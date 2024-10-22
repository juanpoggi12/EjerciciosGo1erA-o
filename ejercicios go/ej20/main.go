package main

import (
	"fmt"
)

const (
	fila = 3
	col  = 4
)

var (
	A                               [fila][col]int
	codigo, cantidad, marca, precio int
)

func main() {

	for f := 0; f < fila; f++ {

		fmt.Println("ingrese el codigo del articulo")
		fmt.Scanln(&codigo)
		fmt.Println("ingrese el codigo de la marca del articulo")
		fmt.Scanln(&marca)
		fmt.Println("ingrese la cantidad existente")
		fmt.Scanln(&cantidad)
		fmt.Println("ingrese el precio de venta del articulo")
		fmt.Scanln(&precio)

		A[f][0] = codigo
		A[f][1] = marca
		A[f][2] = cantidad
		A[f][3] = precio

	}

	fmt.Println("la matriz A es: ", A)

	res1 := mayor50(A)
	fmt.Println("el/los codigos y precio de los productos que poseen una cantidad mayor a 50 son (los datos se encuentran en cada fila de la matriz): ", res1)

	res2 := n2(A)
	fmt.Println("la matriz con los codigos de articulo y de marca de los que poseen precio menor a 300 es: ", res2)

	res3 := ordenar(&A, 3)
	fmt.Println("la matriz ordenada por precios de manera ascendente es: ", res3)

} //main

func mayor50(A [fila][col]int) [][2]int {
	var vec [2]int
	var matriz [][2]int

	for f := 0; f < fila; f++ {

		if A[f][2] > 50 {
			for x := 0; x < col; x++ {
				if x == 0 {
					vec[0] = A[f][0]
				} else {
					if x == 3 {
						vec[1] = A[f][3]
					}
				}

			}
			matriz = append(matriz, vec)

		}

	}

	return matriz
}

func n2(A [fila][col]int) [][2]int {
	var matriz [][2]int
	var vec [2]int

	for f := 0; f < fila; f++ {

		if A[f][3] <= 300 {
			for x := 0; x < 2; x++ {
				if x == 0 {
					vec[0] = A[f][0]
				} else {
					vec[1] = A[f][1]
				}
			}
			matriz = append(matriz, vec)
		}

	}
	return matriz
}

func ordenar(A *[fila][col]int, num int) *[fila][col]int {

	for x1 := 0; x1 < fila-1; x1++ {
		for x2 := x1 + 1; x2 < fila; x2++ {
			if A[x1][num] > A[x2][num] {
				for x := 0; x < col; x++ {
					aux := A[x1][x]
					A[x1][x] = A[x2][x]
					A[x2][x] = aux

				}
			}
		}
	}
	return A
}
