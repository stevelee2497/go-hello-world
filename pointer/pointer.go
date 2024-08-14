// Golang program to illustrate the
// concept of the Pointer to struct
package main

import "fmt"
import "rsc.io/quote"

// taking a structure
type Employee struct {

	// taking variables
	name string
	empid int
}

// Main Function
func main() {

	// creating the instance of the
	// Employee struct type
	emp := Employee{"ABC", 19078}

	// Here, it is the pointer to the struct
	pts := &emp

	// displaying the values
	fmt.Println(pts)

	// updating the value of name
	pts.name = "XYZ"
	emp.name = "3123"
	changeValue(emp)

	fmt.Println(pts)
	fmt.Println(emp)
	fmt.Println(quote.Go())
}

func changeValue(e Employee) {
	e.name = "changed"
}