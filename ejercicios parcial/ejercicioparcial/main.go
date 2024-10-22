package main

import (
	"fmt"
)

const (
	col = 6
)

var (
	DATOS                                                                                              [][col]int
	vec1                                                                                               [col]int
	RES                                                                                                [6][7]int
	SEL                                                                                                [][3]int
	nroservicio, cliente, claseanimal, codigoservicio, importe, atencion, claseanimales, totalanimales int
)

func main() {
	fmt.Println("ingrese el codigo de clase de animal")
	fmt.Scanln(&claseanimal)

	nroservicio = 999

	for claseanimal != 0 {

		fmt.Println("ingrese el codigo del cliente")
		fmt.Scanln(&cliente)
		fmt.Println("ingrese el codigo de servicio")
		fmt.Scanln(&codigoservicio)
		fmt.Println("ingrese el importe cobrado")
		fmt.Scanln(&importe)
		fmt.Println("ingrese si es atencion a domicilio o no")
		fmt.Scanln(&atencion)

		vec1[0] = nroservicio + 1
		vec1[1] = cliente
		vec1[2] = claseanimal
		vec1[3] = codigoservicio
		vec1[4] = importe
		vec1[5] = atencion

		DATOS = append(DATOS, vec1)

		//carga de res
		RES[codigoservicio-1][claseanimal-1] = RES[codigoservicio-1][claseanimal-1] + 1

		fmt.Println("ingrese el codigo de clase de animal")
		fmt.Scanln(&claseanimal)
	} //fin de mientras

	//totales matriz RES
	c(&RES)
	fmt.Println("ingrese la clase de animal que quiere buscar")
	fmt.Scanln(&claseanimales)
	res2 := d(RES, claseanimales)
	fmt.Println("el nombre de servicio que mas se atendio en la clase de animal", claseanimales, "es: ", res2)

	e(RES) //funcion e

	SEL := f(DATOS) //funcion punto f
	fmt.Println("la matriz con los datos de loos que no fueron atencion a domicilio es: ", SEL)

} //main

func c(RES *[6][7]int) {
	var suma int
	for f := 0; f < 6-1; f++ {
		for c := 0; c < 7-1; c++ {
			suma = suma + RES[f][c]
		}
		RES[f][6] = suma
		suma = 0
	}

	for c := 0; c < 7-1; c++ {
		for f := 0; f < 6-1; f++ {
			suma = suma + RES[f][c]
		}
		RES[5][c] = suma
		suma = 0
	}
	for x := 0; x < 6; x++ {
		totalanimales = totalanimales + RES[5][x]

	}
	RES[5][6] = totalanimales
}

func d(RES [6][7]int, codigo int) string {
	var mayor, posmayor int
	var resultado string
	for f := 0; f < 6; f++ {
		if f == 0 {
			mayor = RES[f][codigo-1]
			posmayor = f + 1
		} else {
			if RES[f][codigo-1] > mayor {
				mayor = RES[f][codigo-1]
				posmayor = f + 1

			}
		}

	}
	switch posmayor {
	case 1:
		{
			resultado = "consulta"
		}
	case 2:
		{
			resultado = "clínico"
		}
	case 3:
		{
			resultado = "preventivo"
		}
	case 4:
		{
			resultado = "etología"
		}
	case 5:
		{
			resultado = "vacunación"
		}
	}
	return resultado
}

func e(RES [6][7]int) {
	var (
		porcentaje int
	)

	for x := 0; x < 6; x++ {
		porcentaje = RES[5][x] * 100 / RES[5][6]
		fmt.Println("el porcentaje que la clase de animal", x+1, "representa sobre el total", "(", RES[5][6], ")", "es: ", porcentaje)
	}
}
func f(DATOS [][col]int) [][3]int {
	var vec [3]int
	for f := 0; f < len(DATOS); f++ {
		if DATOS[f][5] == 2 {
			vec[0] = DATOS[f][0]
			vec[1] = DATOS[f][1]
			vec[2] = DATOS[f][4]
			SEL = append(SEL, vec)
		}
	}
	return SEL
}

func g(SEL *[][3]int)
