package bankholidays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuild(t *testing.T) {
	got, err := Build()
	require.NoError(t, err)

	holidays, err := Load()
	require.NoError(t, err)

	want := &Service{
		Holidays: Holidays(holidays),
	}

	require.Equal(t, want, got)

	require.Equal(t, 332, len(holidays))
}
