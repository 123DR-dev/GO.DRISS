package main

import "fmt"

func main() {
	var n, max,i, position int
	max = 0
	i = 0
	n = 1
	for  n != 0 {
		fmt.Println("entrez un nombre")
		fmt.Scanln(&n)
		if i == 1 || n > max {
			max = n
			position = n
		}
	}
	fmt.Println("le nombre max est :", max, "la position est :", position,)
}
