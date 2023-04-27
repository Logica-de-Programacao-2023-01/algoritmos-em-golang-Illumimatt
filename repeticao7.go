package main

import "fmt"

//Faça um algoritmo
//que leia um número inteiro positivo e mostre todos os seus divisores.

func main() {
	var valor int
	fmt.Print("Digite um numero a ser descoberto os divisores: ")
	fmt.Scan(&valor)
	if valor < 0 {
		fmt.Print("Erro")
	}
	fmt.Print("Os divisores de ", valor, " são: ")
	for num := 1; num <= valor; num++ {
		if valor%num == 0 {
			fmt.Println(num)
		}
	}
}
