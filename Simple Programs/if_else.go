package main

import "fmt"

func main() {
	a := 10
	b := 20
	fmt.Printf("value of a is : %d\n", a)
	fmt.Printf("value of b is : %d\n", b)

	if a > b {
		fmt.Printf("a is the greater number\n")
	} else {
		fmt.Printf("b is the greater number\n")
	}
}
