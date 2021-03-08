package service

import "github.com/sancroth/mux-api/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (dcs DefaultCustomerService) GetAllCustomer() ([]domain.Customer,error) {
	return dcs.repo.FindAll()
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}