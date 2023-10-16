package handlers

import (
	"log"
	"net/http"
	"strconv"
	"working/data"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// handle get request
// 	if r.Method == http.MethodGet {
// 		p.getProducts(w, r)
// 		return
// 	}
// 	// handle post request
// 	if r.Method == http.MethodPost {

// 		p.l.Println("POST", r.URL.Path)
// 		p.addProduct(w, r)
// 		return
// 	}
// 	//handle put request
// 	if r.Method == http.MethodPut {
// 		p.updateProduct(0, w, r)
// 		println("PUT", r.URL.Path)
// 		// expect the id in the URI
// 		p.l.Println("PUT", r.URL.Path)
// 		reg := regexp.MustCompile(`/([0-9]+)`)
// 		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
// 		if len(g) != 1 {
// 			p.l.Println("Invalid URI more than one id")
// 			http.Error(w, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}
// 		if len(g[0]) != 2 {
// 			p.l.Println(g)
// 			p.l.Println("Invalid URI more than one capture group")
// 			http.Error(w, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		idString := g[0][1]
// 		id, err := strconv.Atoi(idString)
// 		if err != nil {
// 			p.l.Println("Invalid URI unable to convert to number")
// 			http.Error(w, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}
// 		p.l.Println("got id", id)
// 		// catch all
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 	}
// }

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

	// Below is a pointer to a Product struct that is empty
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}
func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product")
	p.l.Println("got id", id)

	// Below is a pointer to a Product struct that is empty
	prod := &data.Product{}

	errMarshal := prod.FromJSON(r.Body)
	if errMarshal != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	data.UpdateProduct(id, prod)
}
