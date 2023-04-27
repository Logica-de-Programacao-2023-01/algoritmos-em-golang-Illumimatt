package main

import "fmt"

//Faça um algoritmo que leia a idade de uma pessoa e mostre a sua classificação,
//de acordo com a seguinte tabela:
//até 9 anos: mirim
//10 a 13 anos: infantil
//14 a 17 anos: juvenil
//maiores de 18 anos: adulto

func main() {
	var idade int
	fmt.Print("digite sua idade: ")
	fmt.Scan(&idade)
	if idade <= 9 {
		fmt.Println("Mirim")
	} else if idade <= 13 {
		fmt.Println("Infantil")
	} else if idade <= 17 {
		fmt.Println("Juvenil")
	} else {
		fmt.Println("Adulto")
	}
}
