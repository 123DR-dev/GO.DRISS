package main

import "fmt"

func main() {
	var n, max, position, stop int
	max = 0
	for i:=1;i<11;i++ {
		fmt.Println("entrez un nombre")
	    fmt.Scanln(&n)
		if i == 1 || n > max {
		max = n
		stop = 0
		position=i}
	}
	fmt.Println("le nombre max est :",max,"la position est :",position,"fin de code",stop)
}