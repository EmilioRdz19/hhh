package main

import "fmt"

//Estructura
type Persona struct {
	nombre string
	edad   int
}

//metodos
func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\nEdad: %d\n", p.nombre, p.edad)
}

func (p *Persona) editNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
}

//Herencia
type Empleado struct {
	Persona
	sueldo float64
}

func main() {
	p1 := Persona{"Jorge", 21}

	p1.imprimir()
	p1.editNombre("Juan")

	p1.imprimir()

	p2 := Persona{
		nombre: "Angel",
		edad:   22,
	}

	p2.imprimir()
	p2.editNombre("Rafael")
	p2.imprimir()

	em1 := Empleado{
		sueldo: 1500,
	}
	em1.nombre = "Pedro"
	em1.edad = 23
	em1.imprimir()
	fmt.Println(em1)

}
