package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func calcularValorTotal(p Producto) float64 {
	return p.Precio * float64(p.Cantidad)
}

func tieneStock(p Producto) bool {
	if p.Cantidad > 0 {
		return true
	} else {
		return false
	}
}

func main() {
	p1 := Producto{Nombre: "Camisa", Precio: 15.50, Cantidad: 10}
	p2 := Producto{Nombre: "Pantalon", Precio: 30.00, Cantidad: 0}
	p3 := Producto{Nombre: "Zapatos", Precio: 50.00, Cantidad: 5}

	productos := []Producto{p1, p2, p3}

	for i := 0; i < len(productos); i++ {
		prod := productos[i]

		fmt.Println("--- Producto ---")
		fmt.Println("Nombre:", prod.Nombre)
		fmt.Println("Precio:", prod.Precio)
		fmt.Println("Cantidad:", prod.Cantidad)
		fmt.Println("Valor total:", calcularValorTotal(prod))

		if tieneStock(prod) {
			fmt.Println("Estado de stock: disponible")
		} else {
			fmt.Println("Estado de stock: sin stock")
		}
		fmt.Println()
	}
}
