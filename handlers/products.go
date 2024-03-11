// ? This package handlers provides HTTP handlers for managing products in the microservices architecture.
package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Manas8803/learn-microservice/product-api/data"
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
	all_products := data.GetProducts()
	json_data, err := json.Marshal(all_products)
	if err != nil {
		http.Error(w, "Internal Server Error : "+err.Error(), http.StatusInternalServerError)
	}

	w.Write(json_data)
}
