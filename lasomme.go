package main

import "fmt"

func main() {
	var n, i,  s int
	s = 0
	fmt.Println("entrez un nombre")
	fmt.Scanln(&n)
	for i:=1;i < n+1;i++ {
		s = s + i
	} 
	fmt.Println(s+i)
	
}