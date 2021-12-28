package mensajes

import "fmt"

func Hola() {
	fmt.Println("Hola")
}

const mensaje = "Hola constante"

func funcionPrivada() {
	fmt.Println("Hola privada")
}

func Imprimir() {
	fmt.Println(mensaje)
	funcionPrivada()
}
