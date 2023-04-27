package main

import (
	"fmt"
)

//Faça um algoritmo que
//imprima a tabuada de multiplicação de 1 a 10 para um número fornecido pelo usuário.

func main() {
	var valor int
	fmt.Print("Digite um numero a ser impresso a tabuada: ")
	fmt.Scan(&valor)
	for num := 0; num <= 10; num++ {
		fmt.Println(num * valor)
	}
}
