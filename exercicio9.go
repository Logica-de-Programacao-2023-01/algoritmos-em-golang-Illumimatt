package main

import "fmt"

func main() {
	var x, y, z int

	fmt.Print("Digite três numeros: ")
	fmt.Scan(&x, &y, &z)

	if x > y && x > z {
		if y > z {
			fmt.Printf("%d > %d > %d", x, y, z)
		} else {
			fmt.Printf("%d > %d > %d", x, z, y)
		}
	} else if y > x && y > z {
		if x > z {
			fmt.Printf("%d > %d > %d", y, x, z)
		} else {
			fmt.Printf("%d > %d > %d", y, z, x)
		}
	} else if z > x && z > y {
		if y > x {
			fmt.Printf(" %d > %d > %d", z, y, x)
		} else {
			fmt.Printf(" %d > %d > %d", z, x, y)
		}
	}

}
