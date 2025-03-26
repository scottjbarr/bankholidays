package bankholidays

import (
	"bytes"
	"encoding/csv"
	"io"

	_ "embed"
)

//go:embed data/holidays.csv
var data string

func Load() ([]Holiday, error) {
	r := bytes.NewBufferString(data)

	return LoadFromReader(r)
}

func LoadFromReader(r io.Reader) ([]Holiday, error) {
	reader := csv.NewReader(r)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// skip the header row
	records = records[1:]

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
