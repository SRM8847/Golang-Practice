// anonymous functions

package main

import "fmt"

func main() {
	f := func(j int) {
		fmt.Println("printing", j, "from indside an anonymous function")
	}
	for i := range 5 {
		f(i)
	}
}
