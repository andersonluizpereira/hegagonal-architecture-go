# hegagonal-architecture-go

#Generated mocks with mockgen in the interface
mockgen -destination=application/mocks/application.go -source=application/product.go application

#Runnig coverage
go test $(go list ./... | grep -v /./) -coverprofile .testCoverage.txt