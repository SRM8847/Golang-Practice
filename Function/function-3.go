// Multiple return values in a single function
package main

import (
	"errors"
	"fmt"
	"os"
)

func divAndRemainder(num, denum int) (int, int, error) {
	if denum == 0 {
		return 0, 0, errors.New("divisible by zero not possible")
	}
	return num / denum, num % denum, nil
}

func main() {
	result, remainder, err := divAndRemainder(8, 0)

	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}
