package main

import (
	"fmt"
	"paquetes/models"
)

func main() {

	p1 := models.Persona{}
	p1.Constructor("Jose", 21)

	fmt.Println(p1)
	fmt.Println(p1.GetNombre())
	p1.SetNombre("Eduardo")

	fmt.Println(p1)

}
