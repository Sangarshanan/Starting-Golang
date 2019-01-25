package main

import "fmt"

//Refer https://golang.org/pkg/fmt/ for all defintions of %g %T and all that

func main() {
	var a, b, c = 3, 4.12, "foo"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	//%T prints out the type
	fmt.Printf("a is of type %T \n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)

	// SHORT HAND DECLARATION ":=" declares and assigns, "=" just assigns
	i := 10
	fmt.Println(i)

}
