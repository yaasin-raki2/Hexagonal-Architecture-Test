package service

import (
	"github.com/yaasin-raki2/banking/domain"
	"github.com/yaasin-raki2/banking/dto"
	"github.com/yaasin-raki2/banking/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	switch status {
	case "active":
		status = "1"
	case "inactive":
		status = "0"
	default:
		status = ""
	}

	customers, err := s.repo.FindAll(status)

	if err != nil {
		return nil, err
	}

	response := make([]dto.CustomerResponse, len(customers))

	for i := range response {
		response[i] = customers[i].ToDto()
	}

	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)

	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
