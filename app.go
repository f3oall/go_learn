package main

import "fmt"

func Scan(max int) int {
	var choice int
	fmt.Printf("Type a number of required action:")
	fmt.Scanf("%d", &choice)
	if choice <= 0 {
		return Scan(max)
	}
	return choice
}
func main() {
	choice := Scan(10)
	fmt.Println(choice)
}
