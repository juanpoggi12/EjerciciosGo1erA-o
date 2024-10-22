package main

import "fmt"

const (
	mesas = 5
	filam = 5
	colm  = 6
)

var (
	m                                      [filam][colm]int
	partido1, partido2, partido3, partido4 int
)

func main() {

	for c := 0; c < colm-1; c++ {

		fmt.Println("ingrese la cantidad de votos del partido 1")
		fmt.Scanln(&partido1)
		fmt.Println("ingrese la cantidad de votos del partido 2")
		fmt.Scanln(&partido2)
		fmt.Println("ingrese la cantidad de votos del partido 3")
		fmt.Scanln(&partido3)
		fmt.Println("ingrese la cantidad de votos del partido 4")
		fmt.Scanln(&partido4)

		m[0][c] = m[0][c] + partido1
		m[1][c] = m[1][c] + partido2
		m[2][c] = m[2][c] + partido3
		m[3][c] = m[3][c] + partido4

	}

	fmt.Println("la matriz m es:", m)

}
