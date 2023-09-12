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
		p.GetProducts(w, r)
		return
	}
	// handle post request
	if r.Method == http.MethodPost {
		p.AddProduct(w, r)
		return
	}
	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "oops", http.StatusInternalServerError)
		return
	}
}
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")
}
