package main

import (
	"fmt"
	"strings"
)

func obtenerGanadora(votos map[string]int) string {
	var ganadora string
	maxVotos := -1

	for actividad, total := range votos {
		if total > maxVotos {
			maxVotos = total
			ganadora = actividad
		}
	}
	return ganadora
}

func main() {
	votos := map[string]int{
		"deportes":  0,
		"teatro":    0,
		"talleres":  0,
		"excursión": 0,
		"lectura":   0,
	}

	fmt.Println("=== REGISTRO DE VOTOS PARA ACTIVIDAD DE INTEGRACIÓN ===")
	fmt.Println("Actividades disponibles: deportes, teatro, talleres, excursión, lectura")
	fmt.Println()

	var i = 1
	for i <= 5 {
		var eleccion string
		fmt.Println("Ingrese el voto", i, ":")
		fmt.Scan(&eleccion)

		eleccion = strings.ToLower(eleccion)

		if _, existe := votos[eleccion]; existe {
			votos[eleccion]++
			i++
		} else {
			fmt.Println("  -> Actividad no válida. Escriba: deportes, teatro, talleres, excursión o lectura.")
		}
	}

	fmt.Println()
	fmt.Println("--- RESULTADOS FINALES DE LA VOTACIÓN ---")
	for actividad, cantidad := range votos {
		fmt.Println("•", actividad, ":", cantidad, "votos")
	}

	actividadGanadora := obtenerGanadora(votos)
	fmt.Println("La actividad con mayor aceptación es:", actividadGanadora)
}
