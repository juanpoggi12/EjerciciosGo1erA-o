package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	filaPartido = 4
	colPartido  = 5
)

var (
	equipo    [8]string
	nombre    string
	partidos  [filaPartido][colPartido]int
	vec       [2]int
	aleatorio int
)

func main() {
	for f := 0; f < 8; f++ {
		fmt.Println("ingrese el nombre del equipo")
		fmt.Scanln(&nombre)
		equipo[f] = nombre
	}
	fmt.Println("el vector equipo es: ", equipo)
	/////carga de vector equipo 8 elementos

	rand.Seed(time.Now().Unix())

	for f := 0; f < filaPartido; f++ {
		for c := 0; c < colPartido; c++ {
			if c == 0 {
				partidos[f][c] = f + 1
			} else if c == 1 {
				aleatorio = rand.Intn(7) + 1
				vec[c-1] = aleatorio
				partidos[f][c] = aleatorio
			} else if c == 2 {
				aleatorio = rand.Intn(7) + 1
				if aleatorio != vec[c-2] {
					partidos[f][c] = aleatorio
				} else {
					partidos[f][c] = rand.Intn(7) + 1
				}
			}
		}
	}

	fmt.Println(partidos)
}

