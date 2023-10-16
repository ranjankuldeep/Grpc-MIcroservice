package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w http.ResponseWriter) error {
	e := json.NewEncoder(w)
	//This below will Encode method will mutate the p pointer receiver to JSON and write to the ResponseWriter interface w passed in as a parameter to the ToJSON method.
	return e.Encode(p)
}
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
func GetProducts() Products {
	return productList
}
func AddProduct(p *Product) {
	p.ID = getNextId()
	productList = append(productList, p)
}
func UpdateProduct(id int, p *Product) {
	for i := 0; i < len(productList); i++ {
		if id == productList[i].ID {
			productList[i] = p
			fmt.Println("hey product matched")
		}
	}
}
func getNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// create a slice of products
var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
}
