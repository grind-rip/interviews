package main

import "fmt"

// Complete the 'fizzBuzz' function below.
//
// The function accepts INTEGER n as parameter.

// fizzBuzz prints "Fizz" instead of the number for each multiple of 3, prints
// "Buzz" instead of the number for each multiple of 5, and prints "FizzBuzz"
// instead of the number for numbers which are multiples of both 3 and 5.
func fizzBuzz(n int32) {
	// Constraints:
	//   * 0 < n < 2 * 10^5
	if n <= 0 || n >= 200000 {
		return
	}

	for i := int32(1); i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0: // Can also use i%15
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Printf("%d\n", i)
		}
	}
}
