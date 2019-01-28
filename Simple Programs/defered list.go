//A defer statement is often used with paired operations like open and close,
//connect and disconnect, or lock and unlock to ensure that
//resources are released in all cases, no matter how complex the control flow.
//The right place for a defer statement that releases a resource is immediately
//after the resource has been successfully acquired.

package main

import "fmt"

//Go has a special statement called defer that schedules
//a function call to be run after the function completes

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	//Deferred function calls are pushed onto a stack.
	//When a function returns, its deferred calls
	//are executed in last-in-first-out order
	fmt.Println("Printing as LIFO")
}
