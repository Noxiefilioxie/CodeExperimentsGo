package main

import "fmt"

func main() {
	DoFizzBuzz(15, 3, 5)
}

func DoFizzBuzz(n int, f int, b int) {
	for i := 1; i <= n; i++ {

		fizz := i%f == 0
		buzz := i%b == 0
		switch {
		case fizz && buzz:
			fmt.Println("FizzBuzz")
		case fizz:
			fmt.Println("Fizz")
		case buzz:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)

		}

	}
}
