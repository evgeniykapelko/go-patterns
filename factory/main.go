package main

import (
	"factory/products"
	"fmt"
	"time"
)

func main() {
	//factory := products.Product{}
	//
	//product := factory.New()

	product := products.Product{
		ProductName: "Test name",
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}

	fmt.Println("My product was created at", product.CreateAt.UTC())
}
