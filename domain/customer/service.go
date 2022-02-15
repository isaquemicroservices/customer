package customer

import "github.com/isaqueveraslabs/customer-microservice/configuration/database"

// Service models a service base struct
type Service struct {
	repo ICustomer
}

// GetService retrieves a service type
func GetService(r ICustomer) *Service {
	return &Service{repo: r}
}

// GetcustomerRepository retrieve repository for access to customer data
func GetcustomerRepository(db *database.DBTransaction) ICustomer {
	return New(db)
}
