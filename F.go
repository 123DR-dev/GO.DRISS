package main

import "fmt"

import "math"

func main() {
	var racines, a, b, c, Delta int
	Delta = b*b-4*a*c 
	Println("Delta =", Delta)
	if Delta > 0 {
		Println("Delta est positive")
		Println("donc f a deux racines")
		Println("x1=",(-b-sqrt(d))/(2*a))
		Println("  et x2 =",(-b+sqrt(d))/(2*a))
	}else if Delta == 0 {
		
		Println("Delta est nul")
		Println("Donc f a une racine")
		Println("x1=",-b/2*a)
	}else {
		Println("delta est negatif donc f n'a pas de racine")	
	}	
	




	