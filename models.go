package bankholidays

type Holidays []Holiday

type Holiday struct {
	Country string
	Name    string
	Date    string
}

func parseHoliday(code, name, date string) (*Holiday, error) {
	return &Holiday{
		Country: code,
		Name:    name,
		Date:    date,
	}, nil
}

func (h Holidays) Filter(opts *FilterOpts) Holidays {
	return Filter(h, opts)
}
