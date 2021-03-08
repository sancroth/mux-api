package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (crs CustomerRepositoryStub) FindAll() ([]Customer,error) {
	return crs.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			ID:          "1",
			Name:        "Stub_1",
			City:        "Stub_City_1",
			Zipcode:     "Stub_Zip_1",
			DateOfBirth: "Stub_birth_1",
			Status:      "Stub_status_1",
		},
		{
			ID:          "2",
			Name:        "Stub_2",
			City:        "Stub_City_2",
			Zipcode:     "Stub_Zip_2",
			DateOfBirth: "Stub_birth_2",
			Status:      "Stub_status_2",
		},
	}
	return CustomerRepositoryStub{customers: customers}
}