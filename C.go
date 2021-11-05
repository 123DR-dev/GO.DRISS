package main

import "fmt"

func main() {
	var X, Y, Z, T int
	fmt.Printf("donnez le premier nombre")
	fmt.Scanln(&X)
	fmt.Printf("donnez le deuxiéme nombre")
	fmt.Scanln(&Y)
	fmt.Printf("donnez le troisième nombre")
	fmt.Scanln(&Z)
	if Y < X {
		T = X
		X = Y
		Y = T
	}

	if Z < X {
		T = Z
		Z = Y
		Y = X
		X = T
	} else if Z < X {
		T = Z
		Z = Y
		Y = T
	}
	fmt.Println(X, Y, Z)

}
