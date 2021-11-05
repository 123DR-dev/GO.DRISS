package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Printf("Donner le premier nombre : ")
	fmt.Scanln(&x)
	fmt.Printf("Donner le deuxième nombre : ")
	fmt.Scanln(&y)
	fmt.Printf("Donner le troisième nombre : ")
	fmt.Scanln(&z)
	if x < y && y < z {
		fmt.Println("ILs SONT ordonnés")
	} else {
		fmt.Println("Ils NE SONT PAS ordonnés")
	}
}
