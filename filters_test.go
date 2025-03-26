package bankholidays

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var holidays = Holidays{
	{
		Country: "US",
		Name:    "Foo",
		Date:    "2025-03-25",
	},
	{
		Country: "AU",
		Name:    "Foo",
		Date:    "2025-03-26",
	},
	{
		Country: "US",
		Name:    "Bar",
		Date:    "2025-03-26",
	},
	{
		Country: "AU",
		Name:    "Bar",
		Date:    "2025-03-27",
	},
}

func TestFilter(t *testing.T) {
	locBrisbane, err := time.LoadLocation("Australia/Brisbane")
	if err != nil {
		t.Fatal(err)
	}

	locNY, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		opts *FilterOpts
		want Holidays
	}{
		{
			name: "no_filter",
			want: holidays,
		},
		{
			name: "zero_filter",
			opts: &FilterOpts{},
			want: holidays,
		},
		{
			name: "au_only",
			opts: &FilterOpts{
				Country: "AU",
			},
			want: []Holiday{
				holidays[1],
				holidays[3],
			},
		},
		{
			name: "exclude_before",
			opts: &FilterOpts{
				ExcludeBefore: mustParseDate("2025-03-26", locBrisbane),
			},
			want: []Holiday{
				holidays[1],
				holidays[2],
				holidays[3],
			},
		},
		{
			name: "exclude_after",
			opts: &FilterOpts{
				ExcludeAfter: mustParseDate("2025-03-26", locNY),
			},
			want: []Holiday{
				holidays[0],
				holidays[1],
				holidays[2],
			},
		},
		{
			name: "range",
			opts: &FilterOpts{
				ExcludeBefore: mustParseDate("2025-03-26", locBrisbane),
				ExcludeAfter:  mustParseDate("2025-03-26", locNY),
			},
			want: []Holiday{
				holidays[1],
				holidays[2],
			},
		},
		{
			name: "all",
			opts: &FilterOpts{
				Country:       "AU",
				ExcludeBefore: mustParseDate("2025-03-26", locBrisbane),
				ExcludeAfter:  mustParseDate("2025-03-26", locNY),
			},
			want: []Holiday{
				holidays[1],
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(holidays, tt.opts)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFilter_chaining(t *testing.T) {
	locBrisbane, err := time.LoadLocation("Australia/Brisbane")
	if err != nil {
		t.Fatal(err)
	}

	base := Holidays(holidays)

	type filterFunc func() Holidays

	tests := []struct {
		name string
		fn   filterFunc
		want Holidays
	}{
		{
			name: "multiple_empty_filters",
			fn: func() Holidays {
				return base.Filter(nil).Filter(nil)
			},
			want: holidays,
		},
		{
			name: "exclude_before",
			fn: func() Holidays {
				opts := &FilterOpts{
					ExcludeBefore: mustParseDate("2025-03-26", locBrisbane),
				}

				return base.Filter(opts)
			},
			want: Holidays{
				holidays[1],
				holidays[2],
				holidays[3],
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.fn())
		})
	}
}
