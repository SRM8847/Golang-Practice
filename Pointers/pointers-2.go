// pointers in a function, simple representation

package main

import "fmt"

func failupdate(g *int) {
	x := 10
	g = &x
	fmt.Println(g, *g)
}

func main() {
	var a int
	failupdate(&a)
}
