package main

import "fmt"

func main() {
	var r, t, hasil float64
	phi := 3.14

	fmt.Scan(&r, &t)

	hasil = r * r * phi * 1 / 3 * t

	fmt.Print(hasil)

}
