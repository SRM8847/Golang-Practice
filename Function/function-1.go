// // way-1 - with parameters and with return type---------------

// package main

// import "fmt"

// func addition(a, b int) int {
// 	return a + b
// }
// func main() {
// 	sh := addition(2, 3)
// 	fmt.Println(sh)
// }

// // way -2 - with parameters and no return type----------------
// package main

// import "fmt"

// func addition(a, b int) {
// 	var sum int = a + b
// 	fmt.Println(sum)
// }

// func main() {
// 	addition(3, 4)
// }

// //way-3 - with no parameters and no return type---------------

// package main

// import "fmt"

// func addition() {
// 	var a int = 8
// 	var b int = 5
// 	var sum = a + b
// 	fmt.Println(sum)
// }
// func main() {
// 	addition()
// }

//way -4 - with retun type and no parameters--------------

package main

import "fmt"

func addition() int {
	var a int = 9
	var b int = 4
	return a + b
}
func main() {
	s := addition()
	fmt.Println(s)
}
