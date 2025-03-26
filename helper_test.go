package bankholidays

import "time"

func mustParseDate(s string, loc *time.Location) time.Time {
	layout := "2006-01-02"

	if loc == nil {
		loc = time.Now().Location()
	}

	t, err := time.ParseInLocation(layout, s, loc)
	if err != nil {
		panic(err)
	}

	return t
}
