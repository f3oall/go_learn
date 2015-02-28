// Orders
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Order struct { //Order structure, field Items is a slices of structure Good(Type GoodSlc was defined in Goods.go)
	Customer string
	Items    GoodSlc
	Bill     int
	Date     time.Time
}
type OrderSlc []Order

var allOrders OrderSlc

func init() {
	allOrders = initializeOrders()
}
func (orders *OrderSlc) Remove(item int) { //This method removes record from orders slice.
	slice := *orders
	slice = append(slice[:item], slice[item+1:]...)
	*orders = slice
}

func (orders OrderSlc) Save() { //Save data in file(Only for orders)
	//for orders
	file, err := os.OpenFile("orders.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error: %v! File 'orders.json' can't be opened!\n", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	for _, order := range orders {
		err = encoder.Encode(order)
		if err != nil {
			fmt.Println("Error: %v! File 'orders.json' can't be written!\n", err)
		}
	}
}

func initializeOrders() OrderSlc { //Initialize empty slice of Order structures and read data from file in it.
	var orders OrderSlc
	file, err := os.OpenFile("orders.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error: %v! File 'orders.json' can't be opened!\n", err)
		os.Exit(1)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var order Order
	for ; err != io.EOF; orders = append(orders, order) {
		err = decoder.Decode(&order)
		if err != nil && err != io.EOF {
			fmt.Println("Error: %v! File 'orders.json' can't be read!\n", err)
			os.Exit(1)
		}
	}
	return orders
}
func (orders *OrderSlc) Add() { //Add new record to orders.
	var new Order
	fmt.Println("Order adding.\n\n")
	fmt.Println("Choose customer:\n")
	for arrayID, client := range allClients {
		fmt.Printf("%d.%s\n", arrayID+1, client.Name)
	}
	ID := scan(len(allClients))
	new.Customer = allClients[ID-1].Name
	choice := 1
	for choice != 2 {
		fmt.Println("Choose goods:\n")
		for arrayID, good := range allGoods {
			fmt.Printf("%d.%s %d$ %dpc\n", arrayID+1, good.Name, good.Price, good.Amount)
		}
		ID := scan(len(allGoods))
		new.Items = append(new.Items, allGoods[ID-1])
		fmt.Println("Woud you like to add another one good?\n1.Yes.\n2.No")
		fmt.Scanf("%d", &choice)
		fflushStdin()
	}
	for _, item := range new.Items {
		new.Bill = new.Bill + item.Price
	}
	new.Date = time.Now()
	fmt.Println("Adding succesfully.\n")
}
func (orders *OrderSlc) Edit()     {}
func (orders *OrderSlc) Delete() { //Delete record from slice, Remove() method is used there.
	fmt.Println("Order deleting.\n\n")
	fmt.Println("Please choose required order:\n")
	for arrayID, order := range *orders {
		fmt.Printf("%d.Customer: %s, Bill: %d$, Date: %v\n", arrayID+1, order.Customer, order.Bill, order.Date)
	}
	ID := scan(len(*orders))
	orders.Remove(ID - 1)
	fmt.Println("Order has been deleted")
	bufio.NewReader(os.Stdin).ReadBytes('\n') //This string was added because I need pause in the end of the method.
}
func (orders OrderSlc) Show() { //Show list of order.
	fmt.Println("List of Orders.\nn")
	for arrayID, order := range orders {
		fmt.Printf("%d.Customer: %s, Bill: %d$, Date: %v\n", arrayID+1, order.Customer, order.Bill, order.Date)
		fmt.Println("Items:\n")
		for arrayID, item := range order.Items {
			fmt.Printf("%d.%s %d$ %dpc good number: %d\n", arrayID+1, item.Name, item.Price, item.Amount, item.Number)
		}
		fmt.Println("n\n")
	}
	fmt.Printf("There are %d records in table.\n", len(orders))
	bufio.NewReader(os.Stdin).ReadBytes('\n') //This string was added because I need pause in the end of the method.
}
