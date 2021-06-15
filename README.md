# hegagonal-architecture-go

#Generated mocks with mockgen in the interface
mockgen -destination=application/mocks/application.go -source=application/product.go application

#Runnig coverage
go test $(go list ./... | grep -v /./) -coverprofile .testCoverage.txt

#Create database and tables
touch db.sqlite
create table products(id string, name string, price float, status string);

#Running cobra
cobra init  --pkg-name=github.com/acpereira/go-hexagonal

#Test cobra command
#created product
go run main.go cli -a=create -n="Product CLI" -p=25.98
#get product
go run main.go cli -a=get --id 16e6d891-bd3e-4f2f-9c55-5f45fb620cf2