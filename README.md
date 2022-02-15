# Building microservices to manage customer data using Go and gRPC

Command to generate protobuf
```go
$ protoc -I . protos/customer/customer.proto --go_out=plugins=grpc:./application
```

### Create folder for config.json file
```bat
$ sudo mkdir /etc/ms-customer
$ sudo touch /etc/ms-customer/config.json
$ sudo cp ./config.json /etc/ms-customer/config.json
$ sudo chmod 777 /etc/ms-customer/config.json
```
if you changed the config.json file, use the command at the bottom to update the config.json file on your computer
```bat
$ sudo cp ./config.json /etc/ms-customer/config.json
```

command to run the test
```go
$ go test ./... --cover
```

Command to generate test files
```go
$ go test -coverprofile cover.out 
$ go tool cover -html=cover.out -o cover.html
```

Create product table in PostgreSQL 
```sql
```
