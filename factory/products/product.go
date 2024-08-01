package products

import "time"

type Product struct {
	ProductName string    `json:"product_name"`
	CreateAt    time.Time `json:"create_date"`
	UpdateAt    time.Time `json:"update_date"`
}

func (p *Product) New() *Product {
	product := Product{
		//ProductName: p.ProductName,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	return &product
}
