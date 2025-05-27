package domain

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001",
			Name:       "bob",
			City:       "Toronto",
			Zipcode:    "1588",
			DateofBith: "2000-01-01",
			Status:     "1"},
		{Id: "7444",
			Name:       "pedr",
			City:       "Toronto",
			Zipcode:    "9999",
			DateofBith: "2003-01-01",
			Status:     "2"},
	}
	return CustomerRepositoryStub{customers: customers}
}
