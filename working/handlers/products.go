package handlers

import (
	"log"
	"net/http"
	"working/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle get request
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}
	// handle post request
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}
	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "oops", http.StatusInternalServerError)
		return
	}
}
func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	// Below is a pointer to a Product struct that is empty
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}
