package main

import "fmt"

// Faça um algoritmo que leia o salário de um funcionário
// e calcule o seu novo salário com um aumento de 10% se o salário for menor que R$ 1000,00;
// ou de 5% se o salário for maior ou igual a R$ 1000,00.
func main() {
	var salario float64
	fmt.Print("digite seu salario: ")
	fmt.Scan(&salario)

	if salario < 1000 {
		const aumento = 1.1
		resultado := salario * aumento
		fmt.Println(" o seu salario com aumento e: ", resultado)
	} else {
		const aumento = 1.05
		resultado := salario * aumento
		fmt.Println("o seu salario com aumento e: ", resultado)
	}

}
