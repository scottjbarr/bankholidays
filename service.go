package bankholidays

type Service struct {
	Holidays Holidays
}

func Build() (*Service, error) {
	holidays, err := Load()
	if err != nil {
		return nil, err
	}

	service := Service{
		Holidays: Holidays(holidays),
	}

	return &service, nil
}

func (s *Service) All() Holidays {
	return s.Holidays
}

func (s *Service) Filter(opts *FilterOpts) Holidays {
	return s.Holidays.Filter(opts)
}
