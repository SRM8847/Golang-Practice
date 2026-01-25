// pointers: modifying original value via pointer

package main

import "fmt"

// this function actually changes the caller's variable
func successUpdate(g *int) {
	*g = 50 // dereference and assign
}

func main() {

	var a int = 10

	fmt.Println("Before:", a)

	successUpdate(&a)

	fmt.Println("After:", a)
}
