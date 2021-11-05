package main

import "fmt"

func main() {
	var nh float64
	var ph float64= 100 
	if nh > 41 && nh < 45   {
        fmt.Println("veuillez saisir le nombre d'heures",ph*(nh+0.1))
		fmt.Scanln(&nh)

	}
}