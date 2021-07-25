package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {
	findAllSql := "SELECT * FROM customer"

	rows, err := d.client.Query(findAllSql)

	if err != nil {
		log.Fatal(err)
	}

	var customers []Customer

	for rows.Next() {
		var c Customer

		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateofBirth, &c.Status)

		if err != nil {
			log.Fatal(err)
		}

		customers = append(customers, c)
	}

	return customers, nil

}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	connStr := "postgres://postgres:postgres@localhost:5432/banking?sslmode=verify-full"
	client, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)

	return CustomerRepositoryDB{client}
}
