package main

import "fmt"

//Faça um algoritmo que leia a idade de uma pessoa em anos e mostre a idade em dias.

func main() {
	var idade int
	fmt.Print("Insira sua idade em anos: ")
	fmt.Scan(&idade)

	resultado := idade * 365

	fmt.Println("A sua idade em dias é: ", resultado)

}
