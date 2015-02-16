// Goods
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Good struct { //Good structure.
	Name   string
	Number int
	Price  int
	Amount int
}
type GoodSlc []Good //It is needed for  methods, without it methods don't work(Error with same names), i don't understand why.

func (goods *GoodSlc) Remove(item int) { //This method removes record from goods slice.
	slice := *goods
	slice = append(slice[:item], slice[item+1:]...)
	*goods = slice
}

func (goods GoodSlc) Save() { //Save data in file(Only for goods)
	file, err := os.OpenFile("goods.json", os.O_RDWR|os.O_CREATE, 0666) //I don't use O_APPEND, because if i use it i have many copies of data, because i haven't any uniqunes check.
	if err != nil {
		fmt.Println("Error: %v! File 'goods.json' can't be opened!\n", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	for _, good := range goods {
		err = encoder.Encode(good)
		if err != nil {
			fmt.Println("Error: %v! File 'goods.json' can't be written!\n", err)
		}
	}
}

func initializeGoods() GoodSlc { //I think, i must don't repeat myself, so it works like initializeOrders() or initializeClients().
	var goods GoodSlc
	file, err := os.OpenFile("goods.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error: %v! File 'goods.json' can't be opened!\n", err)
		os.Exit(1)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var good Good
	for ; err != io.EOF; goods = append(goods, good) {
		err = decoder.Decode(&good)
		if err != nil && err != io.EOF {
			fmt.Println("Error: %v! File 'goods.json' can't be read!\n", err)
			os.Exit(1)
		}
	}
	return goods
}
func (goods *GoodSlc) Add() { //Method for adding data to goods slice.
	stdin := bufio.NewReader(os.Stdin)
	var new Good
	fmt.Println("Good adding.\n\n")
	fmt.Println("Enter name:")
	fmt.Scanf("%s", &new.Name)
	stdin.ReadString('\n')
	fmt.Println("Enter number:")
	fmt.Scanf("%d", &new.Number)
	stdin.ReadString('\n')
	fmt.Println("Enter price:")
	fmt.Scanf("%d", &new.Price)
	stdin.ReadString('\n')
	fmt.Println("Enter amount:")
	fmt.Scanf("%d", &new.Amount)
	stdin.ReadString('\n')
	*goods = append(*goods, new)
	fmt.Println("Adding succesfully.\n")
}
func (goods GoodSlc) Edit() { //Method for editing data.
	var name string
	fmt.Println("Good editting.\n\n")
	fmt.Println("Please select required good by name:")
	fmt.Scanf("%s", &name)
	fflushStdin()
	for _, good := range goods {
		if good.Name == name {
			fmt.Println("1.Name: ", good.Name, "\n")
			fmt.Println("2.Number:", good.Number, "\n")
			fmt.Println("3.Price:", good.Price, "\n")
			fmt.Println("4.Amount:", good.Amount, "\n")
			fmt.Println("5.Return to main menu.    \n\n")
			choice := scan(5)
			switch choice {
			case 1:
				fmt.Println("Enter name:")
				fmt.Scanf("%s", &good.Name)
				break
			case 2:
				fmt.Println("Enter number:")
				fmt.Scanf("%s", &good.Number)
				break
			case 3:
				fmt.Println("Enter price:")
				fmt.Scanf("%s", &good.Price)
				break
			case 4:
				fmt.Println("Enter amount:")
				fmt.Scanf("%s", &good.Amount)
				break
			case 5:
				return
			}
			fflushStdin()
			fmt.Printf("Good number %s has been eddited.", good.Number)
			break
		}
	}
	fmt.Println("No such good!\n1.Try again.\n2.Return to main menu.")
	choice := scan(2)
	switch choice {
	case 1:
		goods.Edit()
		break
	case 2:
		return
	}
}
func (goods *GoodSlc) Delete() { //Method for deleting data.
	var name string
	fmt.Println("Good deleting.\n\n")
	fmt.Println("Please select required good by name:")
	fmt.Scanf("%s", &name)
	fflushStdin()
	for arrayID, good := range *goods {
		if good.Name == name {
			goods.Remove(arrayID)
			fmt.Println("Good has been deleted!")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			return
		}
	}
	fmt.Println("No such good!\n1.Try again.\n2.Return to main menu.")
	choice := scan(2)
	switch choice {
	case 1:
		goods.Delete()
		break
	case 2:
		return
	}
}
func (goods GoodSlc) Show() { //Method for showing data.
	fmt.Println("List of Goods.\n\n")
	fmt.Println("____________________________________________\n")
	for arrayID, good := range goods {
		fmt.Println("ID:", arrayID, "\n")
		fmt.Println("Name: ", good.Name, "\n")
		fmt.Println("Number:", good.Number, "\n")
		fmt.Println("Price:", good.Price, "\n")
		fmt.Println("Amount:", good.Amount, "\n")
		fmt.Println("Availability:")
		if good.Amount != 0 {
			fmt.Println("available")
		} else {
			fmt.Println("unavailable")
		}
		fmt.Println("____________________________________________n")
	}
	fmt.Printf("There are %d records in table.\n", len(goods))
	bufio.NewReader(os.Stdin).ReadBytes('\n') //Pause. See explain in Orders.Go
}
