package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//Customer ID
type Customer struct {
	CustomerId		int
	CustomerName	string
	SSN				string
}

func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "newuser"
	databasePass := "password"
	databaseName := "crm"
	database, error := sql.Open(databaseDriver, databaseUser + ":" + databasePass + "@/" + databaseName)
	if error != nil {
		panic(error.Error())
	}
	return database
}

func GetCustomers() [] Customer {
	var database *sql.DB

	database = GetConnection()
	defer database.Close()

	var rows *sql.Rows
	var error error
	rows, error = database.Query("SELECT * FROM Customer ORDER BY CustomerId DESC")
	if error != nil {
		panic(error.Error())
	}

	customer := Customer{}
	customers := [] Customer{}

	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		error = rows.Scan(&customerId, &customerName, &ssn)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn

		customers = append(customers, customer)
	}
	return customers
}

func InsertCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()
	defer database.Close()

	var insert *sql.Stmt
	var error error
	insert, error = database.Prepare("INSERT INTO Customer (CustomerId, CustomerName, SSN) VALUES (?,?,?)")
	if error != nil {
		panic(error.Error())
	}
	insert.Exec(customer.CustomerId, customer.CustomerName, customer.SSN)
}

func UpdateCustomer(customer Customer) {

	var database *sql.DB
	database = GetConnection()
	defer database.Close()

	var update *sql.Stmt
	var error error

	update, error = database.Prepare("UPDATE Customer SET CustomerName=?, SSN=? WHERE CustomerId=?")
	if error != nil {
		panic(error.Error())
	}
	update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)
}

func DeleteCustomer(customer Customer) {

	var database *sql.DB
	database = GetConnection()
	defer database.Close()

	var delete *sql.Stmt
	var error error

	delete, error = database.Prepare("DELETE FROM Customer WHERE CustomerId=?")
	if error != nil {
		panic(error.Error())
	}
	delete.Exec(customer.CustomerId)
}

func GetCustomerById(customerId int) Customer {
	var database *sql.DB

	database = GetConnection()
	defer database.Close()

	var rows *sql.Rows
	var error error
	rows, error = database.Query("SELECT * FROM Customer WHERE CustomerId=?", customerId)
	if error != nil {
		panic(error.Error())
	}
	customer := Customer{}
	for rows.Next() {
		var CustomerId int
		var CustomerName string
		var SSN string

		error = rows.Scan(&CustomerId, &CustomerName, &SSN)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = CustomerId
		customer.CustomerName = CustomerName
		customer.SSN = SSN

	}
	return customer
}

