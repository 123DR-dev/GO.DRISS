package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func lireChaine() string {
	var donneeLue string
	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	scanner.Scan()                        // lancement du scanner
	donneeLue = scanner.Text()            // stockage du résultat du scanner dans une variable
	return donneeLue
}
func lireEntier() int {
	var nombre int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nombre, _ = strconv.Atoi(scanner.Text()) // conversion du type string en int
	return nombre
}

func lireReel() float64 {
	var donneeLue string
	var nombreLu float64
	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	scanner.Scan()                        // lancement du scanner
	donneeLue = scanner.Text()            // stockage du résultat du scanner dans une variable
	nombreLu, _ = strconv.ParseFloat(donneeLue, 8)
	return nombreLu
}

func main() {
	var montant, nb200, nb100, nb50, nb20, np10, np5, np2, np1, r int
	fmt.println("veuillez saisir le montant a retirer")
	fmt.Scanln(&montant)
	nb200 = montant / 200
	fmt.println("montant de billet de 200", nb200)
	r = r % 200
	fmt.println("montant de billet de 100", nb100)
	r = r % 100
	fmt.println("montant de billet de 50", nb50)
	r = r % 50
	fmt.println(" montant de billet de 20", nb20)
	r = r % 20
	fmt.println("montant de piece de 10", np10)
	r = r % 10
	fmt.println(" montant de piece de 5", np5)
	r = r % 5
	fmt.println("montant de piece de 2", np2)
	r = r % 2
	fmt.println("montant de piece de 1", np1)
	r = r % 2

}
