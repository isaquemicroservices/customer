package customer

import app "github.com/isaqueveraslabs/customer-microservice/application/customer"

// Server implements proto interface
type Server struct {
	app.UnimplementedCustomerServer
}
