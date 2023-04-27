package main

import "fmt"

//Faça um algoritmo que leia dois números inteiros e mostre o resultado da
//multiplicação entre eles,
//se ambos forem positivos; ou a soma, se pelo menos um deles for negativo.

func main() {
	var x, y float64
	fmt.Print("me de dois numeros")
	fmt.Scan(&x, &y)
	if x < 0 && y < 0 {
		resultado := x * y
		fmt.Println(resultado)
	} else {
		resultado := x + y
		fmt.Println(resultado)
	}
}
