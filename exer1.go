package main

//Faça um algoritmo que leia três números inteiros e mostre a soma entre eles.

import "fmt"

func main() {
	var x1, x2, x3 int
	fmt.Print("Digite seus numeros: ")
	fmt.Scan(&x1, &x2, &x3)

	resultado := x1 + x2 + x3

	fmt.Println("O resultado é: ", resultado)

}
