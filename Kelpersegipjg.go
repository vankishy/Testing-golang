package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Scan(&a, &b)
	fmt.Print(2 * (a + b))
}
