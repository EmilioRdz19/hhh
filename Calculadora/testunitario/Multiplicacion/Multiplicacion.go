package main

import "fmt"

func multiplicar(a, b int) int {
	return a * b
}

func main() {

	//Variables
	var a, b, multiplica int

	//Entrada  de datos
	fmt.Print("Ingrese N01: ")
	fmt.Scanln(&a)

	fmt.Print("Ingrese N02: ")
	fmt.Scanln(&b)

	//Llamar la función multiplicar
	multiplica = multiplicar(a, b)

	//Salida de datos
	fmt.Println("La multiplicacion es:", multiplica)

}
