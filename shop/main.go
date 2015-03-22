package main

import (
	"io"
	"log"
	"net/http"

	"github.com/vtg/flash"
)

//Data is an interface for Products, Clients and Orders types.
type Data interface {
	Append(i Item)
	Edit(id int, i Item)
	Remove(id int)
	List() string
	Save()
	GetName() string
	GetNewItem() Item
	GetItem(id int) Item
	Ask(i Item) map[string]interface{}
	FindByName(name string) int
	Decode(r io.Reader) (Item, error)
}

//Item is an interface for Product,Client and Order structures
type Item interface {
	Show() string
}

func main() {
	r := flash.NewRouter()
	r.Resource("/clients", &ClContr{})
	r.Resource("/products", &PrdContr{})
	r.Resource("/orders", &OrdContr{})
	log.Fatal(http.ListenAndServe(":8080", r))
}
