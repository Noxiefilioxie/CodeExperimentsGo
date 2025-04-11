package main

import (
	"fmt"
)

var array = [5]int{1, 2, 5, 4, 3}

func main() {
	DoBinarySearchRecursion(2, 1, 3)
}

func DoBinarySearchRecursion(target int, start int, end int) {
	if start > end {
		fmt.Println("Not Found")
		return
	}

	middle := end + start/2

	if middle == target {
		println("Target found")
	}
	if array[middle] > target {
		DoBinarySearchRecursion(target, start, middle-1)
	}
	if array[middle] < target {
		DoBinarySearchRecursion(target, middle+1, end)
	}
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
