package main

import (
	"fmt"
	"os"
)

//Interface
type MainMenu interface {
	Add()
	Edit()
	Delete()
	Show()
	Save()
}

//Menus
func Menu(m MainMenu) { // Menu for goods
	clearConsole()
	fmt.Println("1.Add new.\n")
	fmt.Println("2.Edit.\n")
	fmt.Println("3.Delete.\n")
	fmt.Println("4.Show list.\n")
	fmt.Println("5.Return to main menu.\n")
	choice := scan(5)
	action := map[int]func(){
		1: m.Add,
		2: m.Edit,
		3: m.Delete,
		4: m.Show,
		5: m.Save,
	}
	action[choice]()
}

func main() {
	fmt.Println("Hello!")
	for {
		var m MainMenu
		clearConsole()
		fmt.Println("1.Clients.\n")
		fmt.Println("2.Goods.\n")
		fmt.Println("3.Orders.\n")
		fmt.Println("4.Exit.\n")
		choice := scan(4)
		switch choice {
		case 1:
			m = &allClients
		case 2:
			m = &allGoods
		case 3:
			m = &allOrders
		case 4:
			os.Exit(0)
		}
		Menu(m)
	}
}
