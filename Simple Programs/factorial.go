// GENERIC FIND THE FACTORIAL CODE ;P

package main

import "fmt"

func factorial(x int) int {
	fact := 1
	if x == 0 {
		return 1
	} else if x < 0 {
		return 0
	} else {
		for i := 1; i < x+1; i++ {
			fact = fact * i
		}
		return fact
	}
}

func main() {
	fmt.Print("Factorial of: ")
	var n int
	fmt.Scanln(&n)
	fmt.Println("The factorial is", factorial(n))
}
