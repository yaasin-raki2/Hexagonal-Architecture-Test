package domain

import "github.com/yaasin-raki2/banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	ZipCode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(id string) (*Customer, *errs.AppError)
}
