package main

import "fmt"

type Aparatos interface {
	encender() string
}

type iPad struct{}
type Celular struct{}
type Laptop struct{}

func (*iPad) encender() string {
	return "iPad Encender pantalla"
}

func (*Celular) encender() string {
	return "Celular Encender pantalla"
}

func (*Laptop) encender() string {
	return "Laptop Encender pantalla"
}

func encenderAparato(aparatos Aparatos) {
	fmt.Println(aparatos.encender())
}

func main() {
	iPad := iPad{}
	encenderAparato(&iPad)

	Celular := Celular{}
	encenderAparato(&Celular)

	Laptop := Laptop{}
	encenderAparato(&Laptop)
}
