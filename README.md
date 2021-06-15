# hegagonal-architecture-go

#Generated mocks with mockgen in the interface
mockgen -destination=application/mocks/application.go -source=application/product.go application

#Runnig coverage
go test $(go list ./... | grep -v /./) -coverprofile .testCoverage.txt

#Create database and tables
touch db.sqlite
create table products(id string, name string, price float, status string);