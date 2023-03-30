// A simple script, which creates a slice type int
// with numbers from 0 to 10 added with append function.
// Then the program iterates over the slice
// and checks if the iterated number is even or odd.
package main

import "fmt"

func main() {
	numbers := []int{}
	i := 0
	for i <= 10 {
		numbers = append(numbers, i)
		i++
	}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println("The number", number, "is even number")
		} else {
			fmt.Println("The number", number, "is odd number")
		}
	}
}
