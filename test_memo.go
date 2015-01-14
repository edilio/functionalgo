package main

import (
	"fmt"
	"funct"
)

func main() {
	fmt.Println("Testing memoizer")

	m := make(map[int]int)
	m[0] = 0
	m[1] = 1

	fibonacci := funct.Memoizer(m, funct.Fib);

	fmt.Println("===============================================================")
	fmt.Println("Testing fibonacci")
	fmt.Println("5", fibonacci(5))
	fmt.Println("10", fibonacci(10))
	fmt.Println("15", fibonacci(15))
	fmt.Println("25", fibonacci(25))
	fmt.Println("30", fibonacci(30))

	memo := make(map[int]int)
	memo[0] = 1

	factorial := funct.Memoizer(memo, funct.Fact)
	
	fmt.Println("===============================================================")
	fmt.Println("Testing factorial")
	fmt.Println("5", factorial(5))
	fmt.Println("10", factorial(10))
	fmt.Println("15", factorial(15))
	fmt.Println("25", factorial(25))


}
