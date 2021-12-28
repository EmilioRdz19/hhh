package main

import "fmt"

func main() {

	a := [6]int{5, 1, 6, 2, 4, 3} //Valores

	fmt.Println("Array Before Sorting : ", a)

	for i := 0; i < 6; i++ {
		flag := 0 //Tomar la variable bandera
		for j := 0; j < 6-i-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
				flag = 1 //Establecer el indicador como 1, si se produce el intercambio
			}
		}
		if flag == 0 {
			break //Romper el conteno
		}
	}

	fmt.Println("Sorted Array : ", a)
}
