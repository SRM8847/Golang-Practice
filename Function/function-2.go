// functions with structs

package main

import "fmt"

type MyFuncOpts struct {
	firstname string
	lastname  string
	age       int
}

func Myfunc(opts MyFuncOpts) {
	fmt.Println("firstname: ", opts.firstname)
	fmt.Println("lastname: ", opts.lastname)
	fmt.Println("age: ", opts.age)
}
func main() {
	opts := MyFuncOpts{
		firstname: "John",
		lastname:  "rock",
		age:       77,
	}
	Myfunc(opts)
}
