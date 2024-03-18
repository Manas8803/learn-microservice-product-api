// ? This package handlers provides HTTP handlers for managing products in the microservices architecture.
package handlers

import (
	"log"
	"net/http"

	"github.com/Manas8803/learn-microservice-product-api/data"
)

// * Products is a struct representing the handler for managing products.
type Products struct {
	logger *log.Logger
}

// * The NewProducts function takes in a logger (eg : The logger may have preconfigured string as "/products-api") and return a new Products handler.
func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getProducts(w, r)
		return

	case http.MethodPost:
		p.addProducts(w, r)
		return

	}
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	all_products := data.GetProducts()
	err := all_products.ToJson(w)
	if err != nil {
		log.Println("Internal error:", err)
		return
	}
}

func (p *Products) addProducts(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Products")

	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(w, "Unable to decode data", http.StatusBadRequest)
	}
	p.logger.Println("Data : ", prod)
	data.AddProduct(prod)
	all_products := data.GetProducts()
	err = all_products.ToJson(w)
	if err != nil {
		log.Println("Internal error:", err)
		return
	}
	p.logger.Println("Data : ", all_products)
}
