CREATE DATABASE Banking;
-------------------------------------------------------------------------------------------

DROP TABLE IF EXISTS customers CASCADE;

DROP SEQUENCE IF EXISTS customer_id_seq CASCADE;
CREATE SEQUENCE customer_id_seq INCREMENT 1 START 2006;

CREATE TABLE customers (
	customer_id SMALLINT NOT NULL DEFAULT nextval('customer_id_seq') PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	date_of_birth DATE NOT NULL,
	city VARCHAR(100) NOT NULL,
	zipcode VARCHAR(10) NOT NULL,
	status BOOLEAN NOT NULL DEFAULT '1'
);

INSERT INTO customers VALUES
	(2000, 'Steve', '1978-12-15', 'New York, NY', '110075', '1'),
	(2001, 'Arian', '1988-05-21', 'New Burg, NB', '12550', '1'),
	(2002, 'Hadley', '1988-04-30', 'Englewood, NJ', '07631', '1'),
	(2003, 'Ben', '1988-01-04', 'Manchester, NH', '03102', '0'),
	(2004, 'Nina', '1988-05-14', 'Clarkston, MI', '48348', '1'),
	(2005, 'Osman', '1988-11-08', 'Hyattsville, MD', '20782', '0');

--------------------------------------------------------------------------------------------

DROP TABLE IF EXISTS accounts CASCADE;

DROP SEQUENCE IF EXISTS account_id_seq CASCADE;
CREATE SEQUENCE account_id_seq INCREMENT 1 START 95476;

CREATE TABLE accounts (
	account_id INT NOT NULL DEFAULT nextval('account_id_seq') PRIMARY KEY,
	customer_id SMALLINT NOT NULL REFERENCES customers(customer_id),
	opening_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
	account_type VARCHAR(10) NOT NULL,
	pin VARCHAR(10) NOT NULL,
	status BOOLEAN NOT NULL DEFAULT '1'
);

INSERT INTO accounts VALUES
	(95470, 2000, '2021-08-22 10:20:06', 'Saving', '1075', '1'),
	(95471, 2001, '2021-06-15 10:27:22', 'Saving', '1255', '1'),
	(95472, 2002, '2021-08-09 10:27:22', 'Checking', '0763', '1'),
	(95473, 2000, '2021-06-03 10:27:22', 'Saving', '0310', '1'),
	(95474, 2004, '2021-02-27 10:27:22', 'Checking', '4834', '1'),
	(95475, 2005, '2021-03-30 10:27:22', 'Checking', '2078', '0');

---------------------------------------------------------------------------------------------

DROP TABLE IF EXISTS transactions CASCADE;

DROP SEQUENCE IF EXISTS transaction_id_seq CASCADE;
CREATE SEQUENCE transaction_id_seq INCREMENT 1 START 1;

CREATE TABLE transactions (
  transaction_id SMALLINT NOT NULL DEFAULT nextval('transaction_id_seq') PRIMARY KEY,
  account_id SMALLINT NOT NULL REFERENCES accounts(account_id),
  amount SMALLINT NOT NULL,
  transaction_type VARCHAR(10) NOT NULL,
  transaction_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(0)
);
