package testunitario

import "testing"

func TestResta(t *testing.T) {
	total := Resta(30, 25)

	if total != 5 {
		t.Errorf("Resultado Incorrecto")
	}
}

func TestMultiplicar(t *testing.T) {
	total := Multiplicar(10, 10)

	if total != 100 {
		t.Errorf("Resultado Incorrecto")
	}
}
