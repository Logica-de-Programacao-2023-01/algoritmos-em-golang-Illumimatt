package main

//Faça um algoritmo que leia o peso e a altura de uma pessoa
//e calcule o seu IMC (Índice de Massa Corporal).

import "fmt"

func main() {
	var peso, altura float64
	fmt.Print("Digite seu peso e sua altura: ")
	fmt.Scan(&peso, &altura)

	resultado := peso / (altura * altura)

	fmt.Println("O seu IMC é: ", resultado)

}
