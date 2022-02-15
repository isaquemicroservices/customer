package postgres

import "github.com/isaqueveraslabs/customer-microservice/configuration/database"

// PGCustomer implements methods for postgres query execution
type PGCustomer struct {
	DB *database.DBTransaction
}
