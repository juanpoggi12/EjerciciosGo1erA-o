package main

import (
	"fmt"
)

const (
	fila = 5
	col  = 6
)

var (
	a                                              [fila][col]int
	nrovendedor, ventas, sumaventas, sumatrimestre int
)

func main() {
	////////////////////////////////////carga de datos en a
	for f := 0; f < fila-1; f++ {
		fmt.Println("ingrese el numero de vendedor")
		fmt.Scanln(&nrovendedor)
		for c := 0; c < col-1; c++ {
			if c == 0 {
				a[f][c] = nrovendedor
			} else if c < 5 {
				fmt.Println("ingrese las ventas para el trimestre ", c)
				fmt.Scanln(&ventas)
				a[f][c] = ventas
				sumaventas = sumaventas + a[f][c]
			}
		}
		a[f][col-1] = sumaventas
		sumaventas = 0
	}
	for c := 1; c < col; c++ {
		for f := 0; f < fila; f++ {
			sumatrimestre = sumatrimestre + a[f][c]
		}
		a[fila-1][c] = sumatrimestre
		sumatrimestre = 0
	}
	//impresion

	fmt.Println("la matriz a es: ", a)
	for f := 0; f < fila-1; f++ {
		fmt.Println("las ventas del vendedor", f+1, "son:", a[f][col-1])
	}
	for c := 1; c < col-1; c++ {
		fmt.Println("las que se realizaron en el trimestre", c, "son:", a[fila-1][c])
	}
	fmt.Println("las ventas que se realiaron a lo largo del año son: ", a[fila-1][col-1])

	///////////////////////////////////////
	ordenar(&a)
	fmt.Println("la matriz a ordenada por ventas es: ", a)
} //main

func ordenar(a *[fila][col]int) {

	for x1 := 0; x1 < fila-2; x1++ {
		for x2 := x1 + 1; x2 < fila-1; x2++ {
			if a[x1][col-1] < a[x2][col-1] {
				for c := 0; c < col; c++ {
					aux := a[x2][c]
					a[x1][c] = a[x2][c]
					a[x2][c] = aux
				}
			}
		}
	}
}
