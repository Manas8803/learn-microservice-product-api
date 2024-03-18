package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"` //* Note here I have completely removed it from the output.
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

// * ToJson method converts all the products that are already existing or all the rows for each product and encodes them into json array that why we have used "Products" instead of "Product".
// * Also encode-decode is faster in comparison to marshal-unmarshal.
func (p *Products) ToJson(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(p)
}

// * FromJson method adds a product to the list of already existing products so it takes "Product" instead of "Products" as a parameter.
func (p *Product) FromJson(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(p)
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   "2024-03-11",
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   "2024-03-11",
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
}

func getNextID() int {
	p := productList[len(productList)-1]
	return p.ID + 1
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}
