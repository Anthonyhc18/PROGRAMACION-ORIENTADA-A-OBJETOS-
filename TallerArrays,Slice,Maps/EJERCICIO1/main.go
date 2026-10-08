package main

import (
	"fmt"
)

/*/EJERCICIO 1/*/
func main() {

	notas := [6][4]float64{
		{92.0, 88.5, 95.0, 91.0},
		{58.0, 62.5, 70.0, 65.0},
		{74.0, 79.0, 83.5, 80.0},
		{99.0, 95.5, 97.0, 100.0},
		{68.0, 71.0, 66.5, 75.0},
		{82.5, 85.0, 88.0, 79.5},
	}

	promediosEstudiantes := make([]float64, 6)
	notasAltas := make([]float64, 6)
	notasBajas := make([]float64, 6)

	var sumaTotalClase float64
	totalNotas := float64(6 * 4)

	for i := 0; i < 6; i++ {
		var sumaEstudiante float64
		alta := notas[i][0]
		baja := notas[i][0]

		for j := 0; j < 4; j++ {
			nota := notas[i][j]
			sumaEstudiante += nota
			sumaTotalClase += nota

			if nota > alta {
				alta = nota
			}
			if nota < baja {
				baja = nota
			}
		}

		promediosEstudiantes[i] = sumaEstudiante / 4.0
		notasAltas[i] = alta
		notasBajas[i] = baja
	}

	promedioGeneral := sumaTotalClase / totalNotas

	fmt.Println("---- ANÁLISIS DE NOTAS DE ESTUDIANTES ----")
	fmt.Println()

	for i := 0; i < 6; i++ {
		fmt.Println("Estudiante", i+1, ":")
		fmt.Println("  * Promedio:      ", promediosEstudiantes[i])
		fmt.Println("  * Nota más alta: ", notasAltas[i])
		fmt.Println("  * Nota más baja: ", notasBajas[i])
		fmt.Println("----------------------------------------")
	}

	fmt.Println("Promedio general de la clase:", promedioGeneral)
}
