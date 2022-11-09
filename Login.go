package main

import "fmt"

func main() {
	var user, pass string
	var iterasi int

	fmt.Scan(&user, &pass)
	iterasi = 0
	for user != "Admin" && pass != "Admin" {
		iterasi = iterasi + 1
		fmt.Scan(&user, &pass)
	}
	fmt.Println("Login Berhasil")
	fmt.Println(iterasi)
}
