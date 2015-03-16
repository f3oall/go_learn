// utils
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
)

//OpenFile .
func OpenFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Error: %v! File '%s' can't be opened!\r\n", err, filename)
		os.Exit(1)
	}
	return file
}

//scan returns integer value which can be from 1 to max
func scan(max int, conn net.Conn, bufc *bufio.Reader) int {
	var choice int
	fmt.Fprintf(conn, "Enter number of required action:")
	fmt.Fscanf(bufc, "%d", &choice)
	fflushStdin(bufc)
	if (int(choice) < 1) || (int(choice) > max) {
		fmt.Fprintf(conn, "No such action!Try again please.\r\n")
		return scan(max, conn, bufc)
	}
	return choice
}

//clearConsole
func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//fflushStdin clears bufer.
func fflushStdin(bufc *bufio.Reader) {
	bufc.ReadString('\n')
}
