package main

import (
	"fmt"
	"os"
)

//Menus
func goodsMenu(goods *GoodSlc) { // Menu for goods
	for {
		clearConsole()
		fmt.Println("1.Add good.\n")
		fmt.Println("2.Edit good.\n")
		fmt.Println("3.Delete good.\n")
		fmt.Println("4.Show good list.\n")
		fmt.Println("5.Return to main menu.\n")
		choice := scan(5)
		switch choice {
		case 1:
			goods.Add()
			break
		case 2:
			goods.Edit()
			break
		case 3:
			goods.Delete()
			break
		case 4:
			goods.Show()
			break
		case 5:
			goods.Save()
			return
		}
	}
}
func clientsMenu(clients *ClientSlc) { //Menu for clients
	for {
		clearConsole()
		fmt.Println("1.Add client.\n")
		fmt.Println("2.Edit client.\n")
		fmt.Println("3.Delete client.\n")
		fmt.Println("4.Show client list.\n")
		fmt.Println("5.Return to main menu.\n")
		choice := scan(5)
		switch choice {
		case 1:
			clients.Add()
			break
		case 2:
			clients.Edit()
			break
		case 3:
			clients.Delete()
			break
		case 4:
			clients.Show()
			break
		case 5:
			clients.Save()
			return
		}
	}
}
func ordersMenu(orders *OrderSlc, goods *GoodSlc, clients *ClientSlc) { //Menu for orders
	for {
		clearConsole()
		fmt.Println("1.Add order.\n")
		fmt.Println("2.Delete order.n")
		fmt.Println("3.Show orders.\n")
		fmt.Println("4.Return to main menu\n")
		choice := scan(4)
		switch choice {
		case 1:
			orders.Add(*clients, *goods)
			break
		case 2:
			orders.Delete()
			break
		case 3:
			orders.Show()
		case 4:
			orders.Save()
			return
		}
	}
}
func body() { //Body function
	clients := initializeClients()
	goods := initializeGoods()
	orders := initializeOrders()
	for {
		clearConsole()
		fmt.Println("1.Clients.\n")
		fmt.Println("2.Goods.\n")
		fmt.Println("3.Orders.\n")
		fmt.Println("4.Exit.\n")
		choice := scan(4)
		switch choice {
		case 1:
			clientsMenu(&clients)
			break
		case 2:
			goodsMenu(&goods)
			break
		case 3:
			ordersMenu(&orders, &goods, &clients)
			break
		case 4:
			os.Exit(0)
		}
	}

}

func main() {
	body()
}
