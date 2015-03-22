// Orders
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

//Order structure
type Order struct {
	ID          int      `json:"id"`
	Customer    string   `json:"customer"`
	Items       Products `json:"items:"`
	Bill        string   `json:"bill"`
	ItemsAmount int      `json:"items_amount"`
}

//Orders variable is the slice of Order structures
type Orders []Order

var allOrds Orders

func init() {
	allOrds = initOrds()
}

//Read from file into the allOrds slice
func initOrds() Orders {
	var ords Orders
	file := OpenFile("orders.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&ords)
	if err != nil && err != io.EOF {
		fmt.Printf("Error: %v! File 'orders.json' can't be read!\r\n", err)
		os.Exit(1)
	}
	return ords
}

//GetNewItem returns empty object of Order structure
func (ords *Orders) GetNewItem() Item {
	return Order{}
}

//GetItem returns requsted by id object from allOrds slice
func (ords Orders) GetItem(id int) Item {
	return ords[id]
}

//GetName returns string "order"
func (ords *Orders) GetName() string {
	return "order"
}

//Ask returns map with string keys and interface values, where values pass to fmt.Fscanf function and keys are questions which print for user.
func (ords *Orders) Ask(i Item) map[string]interface{} {
	o := i.(Order)
	questions := map[string]interface{}{
		"Enter customer": &o.Customer,
		"Enter goods":    &o.Items[o.ItemsAmount-1].Name,
	}
	return questions
}

//Show returns string which contains data of one order.
func (o Order) Show() string {
	s := "\r\nCustomer: " + o.Customer + "\r\nItems: "
	for _, item := range o.Items {
		s = s + item.Show() + "\r\n____________________________________________\r\n"
	}
	s = s + "Bill: " + o.Bill
	return s
}

//FindByName returns integer value which equal to requsted by name(unlike others this method parameter means Order.Customer) order id.
func (ords Orders) FindByName(name string) int {
	for n, o := range ords {
		if o.Customer == name {
			return n
		}
	}
	return -1
}

//Save data to file
func (ords Orders) Save() {
	file := OpenFile("orders.json")
	encoder := json.NewEncoder(file)
	err := encoder.Encode(ords)
	if err != nil {
		fmt.Printf("Error: %v! File 'orders.json' can't be written!\r\n", err)
	}
	file.Close()
}

//Append required Item to allOrds
func (ords *Orders) Append(i Item) {
	o := i.(Order)
	o.ID = len(*ords)
	for _, g := range allPrds {
		if g.Name == o.Items[o.ItemsAmount-1].Name {
			o.Items[o.ItemsAmount-1] = g
			o.ItemsAmount++
		}
	}
	for _, item := range o.Items {
		o.Bill = o.Bill + item.Price
	}
	*ords = append(*ords, o)
}

//Edit required Item and replaces the old value to the new
func (ords Orders) Edit(id int, i Item) {
	o := i.(Order)
	o.ID = id
	for _, g := range allPrds {
		if g.Name == o.Items[o.ItemsAmount-1].Name {
			o.Items[o.ItemsAmount-1] = g
			o.ItemsAmount++
		}
	}
	for _, item := range o.Items {
		o.Bill = o.Bill + item.Price
	}
	ords[id] = o
}

//Remove required object from slice
func (ords *Orders) Remove(id int) {
	slice := *ords
	slice = append(slice[:id], slice[id+1:]...)
	*ords = slice
}

//List returns string which contains data of allOrds slice.
func (ords Orders) List() string {
	s := "List of Orders.\r\n\r\n"
	for _, o := range ords {
		s = s + o.Show() + "____________________________________________\r\n"
	}
	return s
}

//Decode s...
func (ords *Orders) Decode(r io.Reader) (Item, error) {
	var o Order
	err := json.NewDecoder(r).Decode(&o)
	if err != nil {
		return Order{}, err
	}
	return o, nil
}
