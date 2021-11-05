package main

import "fmt"

func main() {
	var n, i int
	fmt.Println("entrez un nombre")
		fmt.Scanln(&n)
	for i :=1;i<11;i++ {
         n = n + 1
	}
		fmt.Println(n + i)
	}

