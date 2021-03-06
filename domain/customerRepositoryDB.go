package domain

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/yaasin-raki2/banking/errs"
	"github.com/yaasin-raki2/banking/logger"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var customers []Customer
	var err error

	if status == "" {
		err = d.client.Select(&customers, "SELECT * FROM customers")
	} else {
		err = d.client.Select(&customers, "SELECT * FROM customers WHERE status = $1", status)
	}

	if err != nil {
		logger.Error("Error while querying cutomers table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "SELECT * FROM customers WHERE customer_id = $1"

	var c Customer

	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scanning the cutomer " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbAddress, dbPort, dbName,
	)

	client, err := sqlx.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)

	return CustomerRepositoryDB{client}
}
