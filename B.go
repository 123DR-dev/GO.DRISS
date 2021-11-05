package main

import "fmt"

func main() {
	var n float64
	fmt.Println("veuillez saisier la note")
	fmt.Scanln(&n)
	switch{
	case n >= 0 && n < 10:
	fmt.Println(" echec ")
	 case n >= 10 && n < 12:
		fmt.Println("passable")
	case n >= 12 && n < 14:
		fmt.Println("assez bien ")
	case n >= 14 && n < 16:
		fmt.Println("bien ")
		case n >= 16 && n < 20:
			fmt.Println("T.bien ")
    default :
			fmt.Println("Valeur invalide ")
			
			}		
		}

	

		

		

		
	
		

	

		
	