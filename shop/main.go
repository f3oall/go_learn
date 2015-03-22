package main

import (
	"io"
	"log"
	"net/http"
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

/*
//Add function provides communication with user in case when he wants to create new record.
func Add(d Data, conn net.Conn, bufc *bufio.Reader) {
	fmt.Fprintf(conn, "Add %s.\r\n\r\n", d.GetName())
	i := d.GetNewItem()
	questions := d.Ask(i)
	for out, value := range questions {
		fmt.Fprintf(conn, out)
		fmt.Fscanf(bufc, "%s", value)
		fflushStdin(bufc)
	}
	d.Append(i)
	fmt.Fprintf(conn, "Adding succesfully.\r\n")
	fflushStdin(bufc)
}

//Change function provide communication with user  in case when he wants to edit existing record.
func Change(d Data, conn net.Conn, bufc *bufio.Reader) {
	var name string
	fmt.Fprintf(conn, "Edit %s.\r\n\r\n", d.GetName())
	fmt.Fprintf(conn, "Please selecet required %s by name", d.GetName())
	fmt.Fscanf(bufc, "%s", &name)
	fflushStdin(bufc)
	id := d.FindByName(name)
	i := d.GetItem(id)
	fmt.Fprintf(conn, i.Show())
	fmt.Fprintf(conn, "Type name of required field to continue")
	fmt.Fscanf(bufc, "%s", &name)
	fflushStdin(bufc)
	questions := d.Ask(i)
	for out, value := range questions {
		if ("Enter " + name + ":") == out {
			fmt.Fprintf(conn, out)
			fmt.Fscanf(bufc, "%s", value)
			fflushStdin(bufc)
			d.Edit(id, i)
			fmt.Fprintf(conn, "Editting succesfully.\r\n")
			fflushStdin(bufc)
			return
		}
	}
	fmt.Fprintf(conn, "Error!\r\n")
	fflushStdin(bufc)
}

//Delete function provide communication with user in case when he wants to delete existing record.
func Delete(d Data, conn net.Conn, bufc *bufio.Reader) {
	var name string
	fmt.Fprintf(conn, "Delete %s.\r\n\r\n", d.GetName())
	fmt.Fprintf(conn, "Please select required %s by name", d.GetName())
	fmt.Fscanf(bufc, "%s", &name)
	fflushStdin(bufc)
	id := d.FindByName(name)
	d.Remove(id)
	fmt.Fprintf(conn, "Deleting succesfully.\r\n")
	fflushStdin(bufc)
}

//Menu function provides communication with user in case when he wants to choose required action.
func Menu(d Data, conn net.Conn, bufc *bufio.Reader) {
	clearConsole()
	for {
		fmt.Fprintf(conn, "1.Add new.\r\n")
		fmt.Fprintf(conn, "2.Edit.\r\n")
		fmt.Fprintf(conn, "3.Delete.\r\n")
		fmt.Fprintf(conn, "4.Show list.\r\n")
		fmt.Fprintf(conn, "5.Return to main menu.\r\n")
		choice := scan(5, conn, bufc)
		action := map[int]func(){
			1: func() { Add(d, conn, bufc) },
			2: func() { Change(d, conn, bufc) },
			3: func() { Delete(d, conn, bufc) },
			4: func() { fmt.Fprintf(conn, d.List()) },
			5: func() { d.Save() },
		}
		action[choice]()
		if choice == 5 {
			break
		}
	}
}
*/

func main() {
	r := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
