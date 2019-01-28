package main

//Import Packages and using them in Golang
import (
	"fmt"
	"math"
)

//println() prints a new blank line and then your message.
//printf() provides string formatting similar to the printf function

func main() {
	fmt.Println(math.Pi)
	fmt.Printf("Root of two is %g \n", math.Sqrt(2))
	fmt.Printf("Floor bound is %g \n", math.Floor(10.9))
	fmt.Printf("Ceil bound is %g \n", math.Ceil(10.9))
	fmt.Printf("The Given value is NAN : %t \n", math.IsNaN(10.9))

}
