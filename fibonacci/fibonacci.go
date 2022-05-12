package main

import "fmt"

func main() {

	fmt.Println("Up to how many numbers")

	var num int
	fmt.Scanln(&num)

	fmt.Println("Fibonacci sequence")
	fibonacci(num)

}

func fibonacci(num int) {
	var first int = 0
	var second int = 1

	for i := 0; i < num; i++ {
		var temp int = second

		second = first + second
		first = temp

		fmt.Println(second)

	}

}
