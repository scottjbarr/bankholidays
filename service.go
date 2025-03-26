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
