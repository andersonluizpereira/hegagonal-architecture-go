package main

import (
	"database/sql"
	db2 "github.com/acpereira/go-hexagonal/adapters/db"
	"github.com/acpereira/go-hexagonal/application"

)

func main()  {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productAdapter)
	product, _ := productService.Create("Product Example 2", 300)
	productService.Enable(product)
}
