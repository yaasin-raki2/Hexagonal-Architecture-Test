package domain

import (
	"github.com/yaasin-raki2/banking/dto"
	"github.com/yaasin-raki2/banking/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAsText() string {
	status := "active"
	if c.Status == "false" {
		status = "inactive"
	}
	return status
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(id string) (*Customer, *errs.AppError)
}
