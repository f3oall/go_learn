// Products
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

//Product structure
type Product struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Number string `json:"number"`
	Price  string `json:"price"`
	Amount string `json:"amount"`
}

//Products variable is the slice of Product structures
type Products []Product

var allPrds Products

func init() {
	allPrds = initPrd()
}

//Read from file into the allPrds slice
func initPrd() Products {
	var prds Products
	file := OpenFile("products.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&prds)
	if err != nil && err != io.EOF {
		fmt.Printf("Error: %v! File 'products.json' can't be read!\r\n", err)
		os.Exit(1)
	}
	return prds
}

//GetNewItem returns empty object of Product structure
func (prds *Products) GetNewItem() Item {
	return Product{}
}

//GetItem returns requsted by id object from allPrds slice
func (prds Products) GetItem(id int) Item {
	return prds[id]
}

//GetName returns string "product"
func (prds *Products) GetName() string {
	return "product"
}

//Ask returns map with string keys and interface values, where values pass to fmt.Fscanf function and keys are questions which print for user.
func (prds *Products) Ask(i Item) map[string]interface{} {
	p := i.(Product)
	questions := map[string]interface{}{
		"Enter name:":   &p.Name,
		"Enter price":   &p.Price,
		"Enter amount:": &p.Amount,
		"Enter number":  &p.Number,
	}
	return questions
}

//Show returns string which contains data of one product.
func (p Product) Show() string {
	s := "\r\nName: " + p.Name + "\r\nNumber: " + p.Number + "\r\nPrice: " + p.Price + "\r\nAmount: " + p.Amount
	return s
}

//FindByName returns integer value which equal to requsted by name product id.
func (prds Products) FindByName(name string) int {
	for n, p := range prds {
		if p.Name == name {
			return n
		}
	}
	return -1
}
func (prds Products) FindByID(id int64) *Product {
	prd := allPrds[id]
	return &prd
}

//Save data to file
func (prds Products) Save() {
	file := OpenFile("products.json")
	encoder := json.NewEncoder(file)
	err := encoder.Encode(prds)
	if err != nil {
		fmt.Printf("Error: %v! File 'products.json' can't be written!\r\n", err)
	}
	file.Close()
}

//Append required Item to allPrds
func (prds *Products) Append(i Item) {
	p := i.(Product)
	p.ID = len(*prds)
	*prds = append(*prds, p)
}

//Edit required Item and replaces the old value to the new
func (prds Products) Edit(id int, i Item) {
	p := i.(Product)
	p.ID = id
	prds[id] = p
}

//Remove required object from slice
func (prds *Products) Remove(id int) {
	slice := *prds
	slice = append(slice[:id], slice[id+1:]...)
	*prds = slice
}

//List returns string which contains data of allPrds slice
func (prds Products) List() string {
	s := "List of Products.\r\n\r\n"
	for _, p := range prds {
		s = s + p.Show() + "\r\n____________________________________________\r\n"
	}
	return s
}

//Decode s...
func (prds *Products) Decode(r io.Reader) (Item, error) {
	var p Product
	err := json.NewDecoder(r).Decode(&p)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}
