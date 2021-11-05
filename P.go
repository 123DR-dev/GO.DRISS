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
	var montant, b200, b100, b50, b20, p10, p5, p2, p1, reste int
	fmt.Println("entrez un montant a retirer")
	fmt.Scanln(&montant)
	b200 = montant / 200
	reste = montant % 200
	b100 = reste / 100
	reste = reste % 100
	b50 = reste / 50
	reste = reste % 50
	b20 = reste / 20
	reste = reste % 20
	p10 = reste / 10
	reste = reste % 10
	p5 = reste / 5
	reste = reste % 5
	p2 = reste / 2
	reste = reste % 2
	p1 = reste

	fmt.Println(b200, "billets de 200dh et", b100, "billets de 100dh et", b50, "billets de 50dh et", b20, "billets de 20dh et ",
		p10, "pieces de 10dh et", p5, "pieces de 5dh et", p2, "pieces de 2dh et", p1, "pieces de 1dh ")

}
