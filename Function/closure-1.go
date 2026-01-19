// closures are functions declared inside a function

// // variable capture in go ---------------------
// package main

// import "fmt"

// func main() {
// 	a := 20
// 	f := func() {
// 		fmt.Println(a)
// 		a = 33                      // here a = 33
// 		fmt.Println(a)
// 	}

// 	f()
// 	fmt.Println(a)
// }

// variable shadowing in closures -----------------

package main

import "fmt"

func main() {
	a := 20
	f := func() {
		fmt.Println(a)
		a := 33 // here a := 33
		fmt.Println(a)
	}

	f()
	fmt.Println(a)
}
