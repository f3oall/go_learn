// utils
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func scan(max int) int { //Function for simple input.Parameter 'max' is maximum value,that user can enter in console.

	var choice int
	fmt.Println("Enter number of required action:")
	fmt.Scanf("%d", &choice)
	fflushStdin()
	if (choice < 1) || (choice > max) {
		fmt.Println("No such action!Try again please.\n")
		return scan(max)
	} else {
		return choice
	}
}
func clearConsole() { //Clear console from text.
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func fflushStdin() { //Clear input stream from character '\n'.Used after fmt.Scanf().
	stdin := bufio.NewReader(os.Stdin)
	stdin.ReadString('\n')
}
