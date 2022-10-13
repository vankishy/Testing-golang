package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	fmt.Print((a-b)%2 > 0 || (b-a)%2 > 0)
}
