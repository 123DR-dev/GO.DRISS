package main

import ("fmt")
	
	


func main() {
	var tp float64
	const ph = 2 
	var montant int
	var  n1, n2, n5, n10 int
	var reste, arendre, apayer int
	fmt.Println("entrez le nombre d'heures que vous etes retser",tp)
	fmt.Scanln(&tp)
	apayer = int(tp*ph)
	fmt.Println("le montant a payer est",apayer)
	fmt.Scanln(&montant)
     arendre= montant - apayer
	 fmt.Println("le montant a rendre est",arendre)
	 fmt.Scanln(&arendre)
	n1 = montant / 1
	reste = reste % 1
	n2 = montant / 2
	reste = reste % 2
	n5 = montant / 5
	reste = reste % 5
	n10 = montant / 10
	reste = reste % 10
	fmt.Println(n1,"piece de 1dh",n2,"peiece de 2dh",n5,"piece de 5dh",n10,"piece de 10dh")
	


}