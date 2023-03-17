package main

//Faça um algoritmo que leia um número inteiro e mostre o seu dobro, triplo e quadruplo.

import "fmt"

func main() {
	var x int
	fmt.Print("Qual é o seu número? ")
	fmt.Scan(&x)

	dobro := x * 2
	triplo := x * 3
	quadr := x * 4

	fmt.Println("O resultado é: ", dobro, triplo, quadr)
}
