package main

import "fmt"

func restar(a, b int) int {
	return a - b
}

func main() {

	//Variables
	var a, b, resta int

	//Entrada  de datos
	fmt.Print("Ingrese N01: ")
	fmt.Scanln(&a)

	fmt.Print("Ingrese N02: ")
	fmt.Scanln(&b)

	//Llamar la función restar
	resta = restar(a, b)

	//Salida de datos
	fmt.Println("La resta es:", resta)

}
