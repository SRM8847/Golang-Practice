//pointers in go
package main

import "fmt"

func main() {
	var x int32 = 10
	pointerx := &x

	var pointerz *int

	fmt.Println(x, *pointerx, pointerx, &x, &*pointerx, &pointerx)
	fmt.Println(pointerz)

} 
