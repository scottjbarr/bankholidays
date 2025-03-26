package bankholidays

import (
	"bytes"
	"encoding/csv"

	_ "embed"
)

//go:embed data/holidays.csv
var data string

func Load() ([]Holiday, error) {
	reader := csv.NewReader(bytes.NewBufferString(data))

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	holidays := make([]Holiday, len(records))

	for i, rec := range records {
		h, err := parseHoliday(rec[0], rec[1], rec[2])
		if err != nil {
			return nil, err
		}

		holidays[i] = *h
	}

	return holidays, nil
}
