package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/yaasin-raki2/banking/errs"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {
	findAllSql := "SELECT * FROM customers"

	rows, err := d.client.Query(findAllSql)

	if err != nil {
		log.Fatal(err)
	}

	var customers []Customer

	for rows.Next() {
		var c Customer

		err := rows.Scan(&c.Id, &c.Name, &c.DateofBirth, &c.City, &c.ZipCode, &c.Status)

		if err != nil {
			log.Fatal(err)
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "SELECT * FROM customers WHERE customer_id = $1"

	row := d.client.QueryRow(customerSql, id)

	var c Customer

	err := row.Scan(&c.Id, &c.Name, &c.DateofBirth, &c.City, &c.ZipCode, &c.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("Error while scanning the cutomer " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	connStr := "postgres://postgres:postgres@localhost:4000/postgres?sslmode=disable"

	client, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)

	return CustomerRepositoryDB{client}
}
