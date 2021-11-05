package main

import "fmt"

func main() {
	var j, m, a, jl, ml, al int
	fmt.Println("entrez le jour")
	fmt.Scanln(&j)
	fmt.Println("entrez le mois")
	fmt.Scanln(&m)
	fmt.Println("entrez lannée")
	fmt.Scanln(&a)
	if j == 31 && m == 12 {
		jl = 1
		ml = 1
		al = a + 1
	} else if (j <= 31 && m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10) {
            		
		jl = 1
		ml = m + 1
		al = a
	} else {
		jl = j + 1
		ml = 1
		al = a
	}
	if (j == 30 && m == 4 || m == 6 || m == 9 || m == 11 || j == 29 && m == 2 || j == 28 && a % 4 != 0) {
		jl = 1
		ml = m + 1
		al = a
	} else {
		jl = j + 1
		ml = 1
		al = a
	}
	fmt.Println("le jour du lendemain sera", jl, "le mois du lendemain sera", ml, "lannée du lendemain sera", al)

}
