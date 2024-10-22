package main

import (
	"fmt"
)

const (
	fila = 5
	col  = 6
)

var (
	m                                                    [fila][col]int
	barrio, partido1, partido2, partido3, partido4, suma int
	part                                                 [][2]int
)

func main() {

	for nromesa := 0; nromesa < 5; nromesa++ {
		fmt.Println("ingrese el numero de barrio")
		fmt.Scanln(&barrio)
		fmt.Println("ingrese la cantidad de votos del partido 1")
		fmt.Scanln(&partido1)
		fmt.Println("ingrese la cantidad de votos del partido 2")
		fmt.Scanln(&partido2)
		fmt.Println("ingrese la cantidad de votos del partido 3")
		fmt.Scanln(&partido3)
		fmt.Println("ingrese la cantidad de votos del partido 4")
		fmt.Scanln(&partido4)

		m[0][barrio-1] = m[0][barrio-1] + partido1
		m[1][barrio-1] = m[1][barrio-1] + partido2
		m[2][barrio-1] = m[2][barrio-1] + partido3
		m[3][barrio-1] = m[3][barrio-1] + partido4

	} //carga de votos en matriz m

	////////////////////// carga de totales
	for f := 0; f < fila-1; f++ {
		for c := 0; c < col-1; c++ {
			suma = suma + m[f][c]
		}
		m[f][col-1] = suma
		suma = 0
	}

	for c := 0; c < col; c++ {
		for f := 0; f < fila-1; f++ {
			suma = suma + m[f][c]
		}
		m[fila-1][c] = suma
		suma = 0
	}

	fmt.Println("la matriz m es: ", m)
	////// funcion punto a, impresion:
	for f := 0; f < fila-1; f++ {
		res := a(m, f+1)
		fmt.Println("la cantidad de votos del partido", f+1, "es: ", res)
	}
	///// funcion punto b, impresion:
	for c := 0; c < col-1; c++ {
		res := b(m, c+1)
		fmt.Println("la cantidad de votantes del barrio", c+1, "es: ", res)
	}
	//// funcion punto c:
	part = c(m)
	fmt.Println("la matriz part es: ", part)

	d(part)
	fmt.Println("la matriz part ordenada por votos es: ", part)

} /////////main

func a(m [fila][col]int, nropartido int) int {
	return m[nropartido-1][col-1]
}

func b(m [fila][col]int, nrobarrio int) int {
	return m[fila-1][nrobarrio-1]
}

func c(m [fila][col]int) [][2]int {
	var part [][2]int
	var vec [2]int
	for f := 0; f < fila-1; f++ {
		vec[0] = f + 1
		vec[1] = m[f][col-1]
		part = append(part, vec)
	}
	return part
}

func d(part [][2]int) {

	for x1 := 0; x1 < len(part)-1; x1++ {
		for x2 := x1 + 1; x2 < len(part); x2++ {
			if part[x1][1] > part[x2][1] {
				for c := 0; c < 2; c++ {
					aux := part[x2][c]
					part[x1][c] = part[x2][c]
					part[x2][c] = aux
				}
			}
		}
	}
}
