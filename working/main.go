package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		// r.Body has i/o reader interface
		// use ioutil for reading these streams from i/o interface
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "oops", http.StatusBadRequest)
			return
		}
		log.Printf("Data %s\n", d)

		fmt.Fprintf(w, "hello %s\n", d)
	})
	http.ListenAndServe(":8080", nil)
}
