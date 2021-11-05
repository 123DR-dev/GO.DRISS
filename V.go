package main


import "fmt"


func main() {
	var montanttva,totalHT,totalttc,totalavantremise,prixunitaire, quantité, tauxremises, remises  float64  
	const tauxtva float64 = 0.2
	fmt.Printf("donnez le prix de larticle")
	fmt.Scanln(&prixunitaire)
	fmt.Printf("donnez le nombre d'article")
	fmt.Scanln(&quantité)
	fmt.Printf("donnez le pourcentage de la remise")
	fmt.Scanln(&tauxremises)
	fmt.Println("donnez le tauxtva")
	fmt.Scanln(&montanttva)
	totalavantremise = prixunitaire * quantité
	montanttva = totalHT * tauxtva
	totalttc = totalHT + tauxtva
	fmt.Println("total avant remises:",totalavantremise)
	fmt.Println("remise:",remises)
	fmt.Println("totalHT",totalHT)
	fmt.Println("montanttva",montanttva)
	fmt.Println("le totalttc est :",totalttc,"dh")

}
