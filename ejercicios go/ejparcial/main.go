package main

import "fmt"

const (
	col     = 6
	filares = 6
	colres  = 7
	claseave=4
)

var (
	datos                                                                [][col]int
	vec                                                                  [col]int
	nroservicio, claseanimal, codcliente, codservicio, importe, atencion int
	res                                                                  [filares][colres]int
	codigofuncd int
	sel[][3] int
)

func main() {

	fmt.Println("ingrese el codigo de clase de animal")
	fmt.Scanln(&claseanimal)
	nroservicio = 999
	for claseanimal != 0 {
		nroservicio = nroservicio + 1

		fmt.Println("ingrese el codigo de cliente")
		fmt.Scanln(&codcliente)
		fmt.Println("ingrese el tipo de servicio")
		fmt.Scanln(&codservicio)
		fmt.Println("ingrese el importe cobrado")
		fmt.Scanln(&importe)
		fmt.Println("ingrese si es atencion a domicilio o no")
		fmt.Scanln(&atencion)

		vec[0] = nroservicio
		vec[1] = codcliente
		vec[2] = claseanimal
		vec[3] = codservicio
		vec[4] = importe
		vec[5] = atencion

		datos = append(datos, vec)
		/////// carga de matriz res:

		res[codservicio-1][claseanimal-1] = res[codservicio-1][claseanimal-1] + 1


		////// fin del ciclo
		fmt.Println("ingrese el codigo de clase de animal")
		fmt.Scanln(&claseanimal)
	
	}



	fmt.Println("la matriz datos es:", datos)///borrar al final!!!!!!!!!!!!!

	//func punto c:
	c(&res)

	//func punto d:
	res1:=d(res, claseave)
	fmt.Println("el nombre del servicio que mas se atendio de la clase de animal ave es :", res1)

	e(res)/// func punto e

	//func punto f:
	sel=f(datos)

	//func punto g:
	g(sel)

} //main

func c(res*[filares][colres]int){
	var suma int
	for f:=0; f<filares-1; f++{
		for c:=0; c<colres; c++{
			suma=suma+res[f][c]
		}
		res[f][colres-1]=suma
		suma=0
	}

	for c:=0; c<colres-1; c++{
		for f:=0; f<filares-1; f++{
			suma=suma+res[f][c]
		}
		res[filares-1][c]=suma
		suma=0
	}

}

func d(res[filares][colres]int, claseave int)(string){
var serviciomayor, posicion int
var servicio string
for f:=0; f<filares; f++{
	if f==0{
		serviciomayor=f+1
		posicion=res[f][claseave-1]
	} else if res[f][claseave-1]>posicion{
		serviciomayor=f+1
		posicion=res[f][claseave-1]
	}

}

switch serviciomayor{
	case 1: servicio="consulta"
	case 2: servicio="clinico"
	case 3: servicio="preventivo"
	case 4: servicio="etologia"
	case 5: servicio="vacunacion"
	default: servicio="servicio no registrado" 
	}
	return servicio
}

func e(res[filares][colres]int){
var porcentaje float32
var  claseanimal string
	for c:=0; c<colres; c++{
		porcentaje=float32(res[filares-1][c])*100 / float32(res[filares-1][colres-1])

		switch c{
		case 0: claseanimal="perros"
		case 1: claseanimal="gatos"
		case 2: claseanimal="conejos"
		case 3: claseanimal="aves"
		case 4: claseanimal="roedores"
		case 5: claseanimal="otras especies"
		}

		fmt.Println("el porcentaje de la clase animal", claseanimal, "es: ", porcentaje)
	}

}

func f(datos[][col]int)[][3]int{
	var vec[3]int
	var matriz[][3] int
	for f:=0; f<len(datos); f++{
		if datos[f][col-1]==2{
			vec[0]=datos[f][0]
			vec[1]=datos[f][1]
			vec[2]=datos[f][4]

			matriz = append(matriz, vec)
		}
	}
	return matriz
}

func g(sel[][3]int){

	for x1:=0; x1<len(sel)-1; x1++{
		for x2:=x1+1; x2<len(sel); x2++{
			if sel[x1][2]<sel[x2][2]{
				for c:=0; c<len(sel[0]); c++{
					aux:=sel[x1][c]
					sel[x1][c]=sel[x2][c]
					sel[x2][c]=aux

				}
			}
		}
	}
}



