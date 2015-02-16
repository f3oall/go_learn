// Clients
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Client struct {
	Name         string
	Surname      string
	Login        string
	Password     string
	CreditCard   string
	Street       string
	City         string
	State        string
	Zip          string
	OrdersAmount int //Client structure
}

type ClientSlc []Client //The same reason, needed for creating methods.

func (clients *ClientSlc) Remove(item int) {
	slice := *clients
	slice = append(slice[:item], slice[item+1:]...)
	*clients = slice //Method for remove record from slice.
}
func (clients ClientSlc) Save() {
	file, err := os.OpenFile("clients.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error: %v! File 'clients.json' can't be opened!\n", err)
		os.Exit(1)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	for _, client := range clients {
		err = encoder.Encode(client)
		if err != nil {
			fmt.Println("Error: %v! File 'clients.json' can't be written!\n", err)
		}
	} //Save client data.
}
func initializeClients() ClientSlc {
	var clients ClientSlc
	file, err := os.OpenFile("clients.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error: %v! File 'clients.json' can't be opened!\n", err)
		os.Exit(1)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var client Client
	for ; err != io.EOF; clients = append(clients, client) {
		err = decoder.Decode(&client)
		if err != nil && err != io.EOF {
			fmt.Println("Error: %v! File 'clients.json' can't be read!\n", err)
			os.Exit(1)
		}
	}
	return clients //Initialize clients data.
}
func (clients *ClientSlc) Add() {
	clearConsole()
	stdin := bufio.NewReader(os.Stdin)
	var new Client
	fmt.Println("Client adding.\n\n")
	fmt.Println("Enter name:")
	fmt.Scanf("%s", &new.Name)
	stdin.ReadString('\n')
	fmt.Println("Enter surname:")
	fmt.Scanf("%s", &new.Surname)
	stdin.ReadString('\n')
	fmt.Println("Enter login:")
	fmt.Scanf("%s", &new.Login)
	stdin.ReadString('\n')
	fmt.Println("Enter password:")
	fmt.Scanf("%s", &new.Password)
	stdin.ReadString('\n')
	fmt.Println("Enter credit card:")
	fmt.Scanf("%s", &new.CreditCard)
	stdin.ReadString('\n')
	fmt.Println("Enter street:")
	fmt.Scanf("%s", &new.Street)
	stdin.ReadString('\n')
	fmt.Println("Enter city:")
	fmt.Scanf("%s", &new.City)
	stdin.ReadString('\n')
	fmt.Println("Enter state:")
	fmt.Scanf("%s", &new.State)
	stdin.ReadString('\n')
	fmt.Println("Enter ZIP:")
	fmt.Scanf("%s", &new.Zip)
	stdin.ReadString('\n')
	new.OrdersAmount = 0
	/*slice := *clients
	  slice = append(slice, new)
	  *clients = slice*/
	*clients = append(*clients, new)
	fmt.Println("Adding succesfully.\n") //Add new record.
}
func (clients ClientSlc) Edit() {
	clearConsole()
	var name string
	fmt.Println("Client editting.\n\n")
	fmt.Println("Please select required client by name:")
	fmt.Scanf("%s", &name)
	fflushStdin()
	for ID, client := range clients {
		if client.Name == name {
			fmt.Println("1.Name: ", client.Name, "\n")
			fmt.Println("2.Surame:", client.Surname, "\n")
			fmt.Println("3.Login:", client.Login, "\n")
			fmt.Println("4.Password:", client.Password, "\n")
			fmt.Println("5.Credit card: ", client.CreditCard, "\n")
			fmt.Println("6.Street:", client.Street, "\n")
			fmt.Println("7.City:", client.City, "\n")
			fmt.Println("8.State:", client.State, "\n")
			fmt.Println("9.ZIP:", client.Zip, "\n")
			fmt.Println("10.Return to main menu.\n\n")
			choice := scan(10)
			switch choice {
			case 1:
				fmt.Println("Enter name:")
				fmt.Scanf("%s", &client.Name)
				break
			case 2:
				fmt.Println("Enter surname:")
				fmt.Scanf("%s", &client.Surname)
				break
			case 3:
				fmt.Println("Enter login:")
				fmt.Scanf("%s", &client.Login)
				break
			case 4:
				fmt.Println("Enter password:")
				fmt.Scanf("%s", &client.Password)
				break
			case 5:
				fmt.Println("Enter credit card:")
				fmt.Scanf("%s", &client.CreditCard)
				break
			case 6:
				fmt.Println("Enter street:")
				fmt.Scanf("%s", &client.Street)
				break
			case 7:
				fmt.Println("Enter city:")
				fmt.Scanf("%s", &client.City)
				break
			case 8:
				fmt.Println("Enter state:")
				fmt.Scanf("%s", &client.State)
				break
			case 9:
				fmt.Println("Enter ZIP:")
				fmt.Scanf("%s", &client.Zip)
				break
			case 10:
				return
			}
			clients[ID] = client
			fflushStdin()
			fmt.Println("Client %d has been eddited.")
			return
		}
	}
	fmt.Println("No such client!\n1.Try again.\n2.Return to main menu.")
	choice := scan(2)
	switch choice {
	case 1:
		clients.Edit()
		break
	case 2:
		return
	} //Edit record.
}
func (clients *ClientSlc) Delete() {
	clearConsole()
	var name string
	fmt.Println("Client deleting.\n\n")
	fmt.Println("Please select required client by name:")
	fmt.Scanf("%s", &name)
	fflushStdin()
	for arrayID, client := range *clients {
		if client.Name == name {
			clients.Remove(arrayID)
			fmt.Println("Client has been deleted")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			return
		}
	}
	fmt.Println("No such client!\n1.Try again.\n2.Return to main menu.")
	choice := scan(2)
	switch choice {
	case 1:
		clients.Delete()
		break
	case 2:
		return
	} //Delete record.
}

func (clients ClientSlc) Show() {
	clearConsole()
	fmt.Println("List of Clients.\n\n")
	fmt.Println("____________________________________________\n")
	for arrayID, client := range clients {
		fmt.Println("ID:", arrayID, "\n")
		fmt.Println("Name: ", client.Name, "\n")
		fmt.Println("Surame:", client.Surname, "\n")
		fmt.Println("Login:", client.Login, "\n")
		fmt.Println("Password:", client.Password, "\n")
		fmt.Println("Credit card: ", client.CreditCard, "\n")
		fmt.Println("Street:", client.Street, "\n")
		fmt.Println("City:", client.City, "\n")
		fmt.Println("State:", client.State, "\n")
		fmt.Println("ZIP:", client.Zip, "\n")
		fmt.Println("____________________________________________n")
	}
	fmt.Printf("There are %d records in table.\n", len(clients))
	bufio.NewReader(os.Stdin).ReadBytes('\n') //Show list of clients.
}
