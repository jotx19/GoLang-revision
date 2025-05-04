package main

import "fmt"

func calculation(a int, b int) int {
	return a + b
}

func main() {
	var a, b, c int

	fmt.Println("Enter a number:")
	fmt.Scanf("%d", &a)

	fmt.Println("Enter another number:")
	fmt.Scanf("%d", &b)

	c = calculation(a, b)

	fmt.Println("The sum is:", c)
}
