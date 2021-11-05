package main

import "fmt"

func main() {
	var heure, minute, seconde int

	for i := 0; i < 60; i++ {

		if heure == 2 {
			break
		}

		heure = i
		fmt.Println(heure, "heure", minute, "minute", seconde, "seconde")

		for j := 0; j < 60; j++ {

			minute = j

			if minute == 60 && seconde == 60 {
				fmt.Println(heure, "heure", minute, "minute", seconde, "seconde")
				minute = 0
				heure++

			} else {
				fmt.Println(heure, "heure", minute, "minute", seconde, "seconde")
			}

			for k := 0; k < 60; k++ {

				seconde = k

				if seconde == 60 {
					fmt.Println(heure, "heure", minute, "minute", seconde, "seconde")
					seconde = 0
					minute++

				} else {
					fmt.Println(heure, "heure", minute, "minute", seconde, "seconde")
				}

			}

		}
	}

	
	
	

}