package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	fila = 3

	col = 7
)

var (
	num int
	A   [fila][col]int
)

func main() {
	rand.Seed(time.Now().Unix())
	for f := 0; f < fila; f++ {
		for c := 0; c < col; c++ {
			if c == 0 {
				A[f][c] = f + 1
			} else {
				A[f][c] = rand.Intn(10)
			}
		}
	}

	fmt.Println("la matriz A cargada es:,", A)

	promedioA(A)

	promedioE(A)

	fmt.Println("ingrese el numero de estudiante que quiere buscar")
	fmt.Scanln(&num)
	res1, res2 := x(A, num)
	fmt.Println("la cantidad de materias que desaprobo el alumno numero", num, "es :", res1)
	fmt.Println("y la cantidad de materias que aprobo es: ", res2)
}

func promedioA(A [fila][col]int) {
	var suma int
	var promedio int

	for c := 1; c < col; c++ {
		for f := 0; f < fila; f++ {
			suma = suma + A[f][c]
		}
		promedio = (suma / fila) + 1
		fmt.Println("el promedio de la asignatura numero", c, "es: ", promedio)
		suma = 0
	}
}

func promedioE(A [fila][col]int) {
	var suma, promedio int

	for f := 0; f < fila; f++ {
		for c := 1; c < col; c++ {

			suma = suma + A[f][c]
		}
		promedio = (suma/col - 1) + 1
		fmt.Println("el promedio de el estudiante numero", f+1, "es: ", promedio)
		suma = 0
	}
}

func x(A [fila][col]int, x int) (int, int) {
	var aprobados, desaprobados int

	for c := 1; c < col; c++ {
		if A[x-1][c] >= 6 {
			aprobados = aprobados + 1
		} else {
			desaprobados = desaprobados + 1
		}

	}

	return desaprobados, aprobados
}
