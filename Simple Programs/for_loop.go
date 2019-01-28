package main

import "fmt"

//  SUM OF N NATURAL NUMBERS

func main() {
	var n int
	fmt.Print("Enter the value of N: ")
	fmt.Scan(&n)
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Print("Sum : ", sum, "\n")
}
