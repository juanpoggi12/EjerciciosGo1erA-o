package main

import (
	"fmt"
)

const (
	col = 4
)

var (
	A                                                    [][col]int
	codigolibro, codigoautor, ejemplares, precio, puntoa int
	vec                                                  [4]int
)

func main() {
	fmt.Println("ingrese el codigo de libro")
	fmt.Scanln(&codigolibro)

	for codigolibro != 0 { //carga de datos//
		fmt.Println("ingrese el codigo de autor")
		fmt.Scanln(&codigoautor)
		fmt.Println("ingrese el numero de ejemplares")
		fmt.Scanln(&ejemplares)
		fmt.Println("ingrese el precio del libro")
		fmt.Scanln(&precio)

		vec[0] = codigolibro
		vec[1] = codigoautor
		vec[2] = ejemplares
		vec[3] = precio

		A = append(A, vec)
		fmt.Println("ingrese el codigo de libro")
		fmt.Scanln(&codigolibro)
	}

	//func punto a:

	fmt.Println("ingrese el codigo de autor que quiere analizar en la matriz")
	fmt.Scanln(&puntoa)
	res1 := a(A, puntoa)
	fmt.Println("los codigos de libro que el autor ingresado publico son: ", res1)

	b(A) //func punto b

}

func a(A [][col]int, num int) (vec []int) {
	for f := 0; f < len(A); f++ {
		if A[f][1] == num {
			vec = append(vec, A[f][0])
		}
	}
	return vec
}

func b(A [][col]int) {
	var libro, autor, precio int
	for f := 0; f < len(A); f++ {

		if f == 0 {
			precio = A[f][3]
			libro = A[f][0]
			autor = A[f][1]
		} else {
			if A[f][3] > precio {
				libro = A[f][0]
				autor = A[f][1]
				precio = A[f][3]
			}

		}
	}
	fmt.Println("el precio del libro mas caro es: ", precio, "su codigo de autor es: ", autor, "su codigo de libro es: ", libro)
}
