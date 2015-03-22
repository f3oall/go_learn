// Clients
package main

import (
	"encoding/json"
	"fmt"
	"io"
)

//Client structure
type Client struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Login        string `json:"login"`
	Password     string `json:"password"`
	CreditCard   string `json:"credit_card"`
	Street       string `json:"street"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          string `json:"zip"`
	OrdersAmount int    `json:"orders_amount"`
}

//Clients variable is a slice of Client structures
type Clients []Client

var allCls Clients

func init() {
	allCls = initCls()
}

//Read from file into the allCls slice
func initCls() Clients {
	var cls Clients
	file := OpenFile("clients.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&cls)
	if err != nil && err != io.EOF {
		fmt.Printf("Error: %v! File 'clients.json' can't be read!\r\n", err)
	}
	file.Close()
	return cls
}

//GetNewItem returns empty object of Client structure
func (cls *Clients) GetNewItem() Item {
	return Client{}
}

//GetItem returns requsted by id object from allCls slice
func (cls Clients) GetItem(id int) Item {
	return cls[id]
}

//GetName returns string "client"
func (cls *Clients) GetName() string {
	return "client"
}

//Ask returns map with string keys and interface values, where values pass to fmt.Fscanf function and keys are questions which print for user.
func (cls *Clients) Ask(i Item) map[string]interface{} {
	c := i.(Client)
	questions := map[string]interface{}{
		"Enter name:":        &c.Name,
		"Enter surname:":     &c.Surname,
		"Enter login:":       &c.Login,
		"Enter password:":    &c.Password,
		"Enter street:":      &c.Street,
		"Enter city:":        &c.City,
		"Enter state:":       &c.State,
		"Enter credit card:": &c.CreditCard,
		"Enter ZIP:":         &c.Zip,
	}
	return questions
}

//Show returns string which contains data of one client.
func (c Client) Show() string {
	s := "\r\nName: " + c.Name + "\r\nSurname: " + c.Surname + "\r\nLogin: " + c.Login
	s = s + "\r\nPassword: " + c.Password + "\r\nCredit card: " + c.CreditCard
	s = s + "\r\nStreet: " + c.Street + "\r\nCity: " + c.City + "\r\nState: " + c.State + "\r\nZip: " + c.Zip
	return s
}

//FindByName returns integer value which equal to requsted by name client id.
func (cls Clients) FindByName(name string) int {
	for n, c := range cls {
		if c.Name == name {
			return n
		}
	}
	return -1
}

//Save data to file
func (cls Clients) Save() {
	file := OpenFile("clients.json")
	encoder := json.NewEncoder(file)
	err := encoder.Encode(cls)
	if err != nil {
		fmt.Printf("Error: %v! File 'clients.json' can't be written!\r\n", err)
	}
	file.Close()
}

//Append required Item to allCls
func (cls *Clients) Append(i Item) {
	c := i.(Client)
	c.ID = len(*cls)
	*cls = append(*cls, c)
}

//Edit required Item and replaces the old value to the new
func (cls Clients) Edit(id int, i Item) {
	c := i.(Client)
	c.ID = id
	cls[id] = c
}

//Remove required object from slice
func (cls *Clients) Remove(id int) {
	slice := *cls
	slice = append(slice[:id], slice[id+1:]...)
	*cls = slice
}

//List returns string which contains data of allCls slice.
func (cls Clients) List() string {
	s := "List of Clients.\r\n\r\n"
	for _, c := range cls {
		s = s + c.Show() + "____________________________________________\r\n"
	}
	return s
}

//Decode s...
func (cls *Clients) Decode(r io.Reader) (Item, error) {
	var c Client
	err := json.NewDecoder(r).Decode(&c)
	if err != nil {
		return Client{}, err
	}
	return c, nil
}
