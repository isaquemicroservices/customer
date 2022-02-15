package customer

import (
	"github.com/isaqueveraslabs/customer-microservice/configuration/database"
	"github.com/isaqueveraslabs/customer-microservice/infrastructure/persistence/customer/postgres"
)

// repository is a base structure that implements methods specified by Icustomer
type repository struct {
	pg *postgres.PGCustomer
}

// New creates a new customer repository
func New(db *database.DBTransaction) *repository {
	return &repository{pg: &postgres.PGCustomer{DB: db}}
}
