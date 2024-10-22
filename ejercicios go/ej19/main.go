package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	fila = 10
	col  = 3
)

var (
	A         [fila][col]int
	M         [][col]int
	vectorM   [col]int
	aprobados int
	C         [fila][col + 1]int
)

func main() {
	rand.Seed(time.Now().Unix())

	for f := 0; f < fila; f++ {
		for c := 0; c < col; c++ {
			if c == 0 {
				for x := 0; x < fila; x++ {
					A[x][c] = x + 1
				}
			} else {
				A[f][c] = rand.Intn(10)
			}
		}
	}

	fmt.Println("la matriz A es:, ", A)

	/////// matriz M:
	res1 := matrizM(A)
	fmt.Println("la matriz M es: ", res1)

	res2 := matrizC(A)
	fmt.Println("la matriz C es: ", res2)

	res3 := ordenar(C)
	fmt.Println("la matriz C ordenada ascendentemente segun los promedios es: ", res3)

} //main

func ordenar(C [fila][col + 1]int) [fila][col + 1]int {

	for x1 := 0; x1 < fila-1; x1++ {
		for x2 := x1 + 1; x2 < fila; x2++ {
			if C[x1][col] > C[x2][col] {
				for c := 0; c < col+1; c++ {
					aux := C[x1][c]
					C[x1][c] = C[x2][c]
					C[x2][c] = aux
				}
			}

		}

	}
	return C
}

func matrizC(A [fila][col]int) [fila][col + 1]int {
	var suma int
	for f := 0; f < fila; f++ {
		suma = 0
		for c := 0; c < col; c++ {
			if c == 0 {
				C[f][c] = A[f][c]
			} else {
				suma = suma + A[f][c]
				C[f][c] = A[f][c]
			}
		}
		C[f][col] = suma / (col - 1)

	}
	return C
}

func matrizM(A [fila][col]int) [][col]int {
	var vector [col]int

	for f := 0; f < fila; f++ {
		for c := 1; c < col; c++ {
			if A[f][c] >= 6 {
				aprobados = aprobados + 1
			}
		}
		if aprobados == col-1 {
			for x := 0; x < col; x++ {
				vector[x] = A[f][x]
			}
			M = append(M, vector)
		}
		aprobados = 0
	}
	return M
}
