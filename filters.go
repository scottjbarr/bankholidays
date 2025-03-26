package bankholidays

import "time"

const (
	layoutDate = "2006-01-02"
)

type FilterOpts struct {
	Country       string
	ExcludeBefore time.Time
	ExcludeAfter  time.Time
}

func (f FilterOpts) IsZero() bool {
	return f.Country == "" && f.ExcludeBefore.IsZero() && f.ExcludeAfter.IsZero()
}

// Filter allows filtering a slice of OHLC.
func Filter(in Holidays, opts *FilterOpts) Holidays {
	if opts == nil || opts.IsZero() {
		return in
	}

	out := []Holiday{}

	for _, h := range in {
		if opts.Country != "" && h.Country != opts.Country {
			continue
		}

		if opts.ExcludeBefore.Unix() > 0 && h.Date < dateFormat(opts.ExcludeBefore) {
			continue
		}

		if opts.ExcludeAfter.Unix() > 0 && h.Date > dateFormat(opts.ExcludeAfter) {
			continue
		}

		out = append(out, h)
	}

	return out
}

func dateFormat(t time.Time) string {
	return t.Format(layoutDate)
}
