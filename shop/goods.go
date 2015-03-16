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
	Name   string
	Number string
	Price  string
	Amount string
}

//Products variable is the slice of Product structures
type Products []Product

var allPrds Products

func init() {
	allPrds = initPrd()
}

//Read from file into the allPrds slice
func initPrd() Products {
	var prd Products
	file := OpenFile("products.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&prd)
	if err != nil && err != io.EOF {
		fmt.Printf("Error: %v! File 'products.json' can't be read!\r\n", err)
		os.Exit(1)
	}
	return prd
}

//GetNewItem returns empty object of Product structure
func (prd *Products) GetNewItem() Item {
	return Product{}
}

//GetItem returns requsted by id object from allPrds slice
func (prd Products) GetItem(id int) Item {
	return prd[id]
}

//GetName returns string "product"
func (prd *Products) GetName() string {
	return "product"
}

//Ask returns map with string keys and interface values, where values pass to fmt.Fscanf function and keys are questions which print for user.
func (prd *Products) Ask(i Item) map[string]interface{} {
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
func (prd Products) FindByName(name string) int {
	for n, p := range prd {
		if p.Name == name {
			return n
		}
	}
	return -1
}

//Save data to file
func (prd Products) Save() {
	file := OpenFile("products.json")
	encoder := json.NewEncoder(file)
	err := encoder.Encode(prd)
	if err != nil {
		fmt.Printf("Error: %v! File 'products.json' can't be written!\r\n", err)
	}
	file.Close()
}

//Append required Item to allPrds
func (prd *Products) Append(i Item) {
	p := i.(Product)
	*prd = append(*prd, p)
}

//Edit required Item and replaces the old value to the new
func (prd Products) Edit(id int, i Item) {
	p := i.(Product)
	prd[id] = p
}

//Remove required object from slice
func (prd *Products) Remove(id int) {
	slice := *prd
	slice = append(slice[:id], slice[id+1:]...)
	*prd = slice
}

//List returns string which contains data of allPrds slice
func (prd Products) List() string {
	s := "List of Products.\r\n\r\n"
	for _, p := range prd {
		s = s + p.Show() + "\r\n____________________________________________\r\n"
	}
	return s
}
