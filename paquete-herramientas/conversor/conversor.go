package conversor

import "fmt"

func Convertir() {
	var dolares float64
	var opcion string

	fmt.Println("🤑CONVERSOR DE MONEDAS💰")
	fmt.Print("Ingrese el valor en dólares: ")
	fmt.Scan(&dolares)

	if dolares < 0 {
		fmt.Println("El valor no puede ser negativo.")
		return
	}

	fmt.Println("Monedas disponibles:")
	fmt.Println("1. Euros")
	fmt.Println("2. Libras Esterlinas (LB)")
	fmt.Println("3. Won Surcoreano")
	fmt.Println("4. Bitcoin (BTC)")
	fmt.Print("Seleccione una opción (1-4): ")
	fmt.Scan(&opcion)

	var resultado float64

	switch opcion {
	case "1":
		resultado = dolares * 0.92
		fmt.Println(dolares, "dólares son", resultado, "Euros")
	case "2":
		resultado = dolares * 0.79
		fmt.Println(dolares, "dólares son", resultado, "Libras Esterlinas")
	case "3":
		resultado = dolares * 1330.0
		fmt.Println(dolares, "dólares son", resultado, "Won Surcoreano")
	case "4":
		resultado = dolares * 0.000015
		fmt.Println(dolares, "dólares son", resultado, "Bitcoins")
	default:
		fmt.Println("Opción no válida.")
	}
}
