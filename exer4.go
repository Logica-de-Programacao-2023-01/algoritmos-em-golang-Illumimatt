package main

//Faça um algoritmo que leia três números reais
//e mostre a média ponderada entre eles, com pesos 2, 3 e 5, respectivamente.

import "fmt"

func main() {
	var x, y, z float64
	fmt.Print("Digite seus três numeros: ")
	fmt.Scan(&x, &y, &z)

	resultado := ((x * 2) + (y * 3) + (z * 5)) / 10
	fmt.Println("A sua média é : ", resultado)

}
