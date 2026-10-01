package main

import (
	"fmt"
)

func main() {
	fmt.Println("Bienvenido a este programa. Que desea hacer?")
	fmt.Println("1. Registrar una nueva venta")
	fmt.Println("2. Mostrar estadísticas")
	fmt.Println("3. Salir")
	var usrOpcion int
	subtotal := []float64{}
	nombresProductos := []string{}
	for {
		fmt.Println("Que deseas hacer (Ingresa un numero): ")
		fmt.Scan(&usrOpcion)
		if usrOpcion == 1 {
			var usrOpcionProducto int
			var cantidadVendida float64
			productos := [3]string{"Arroz", "Leche", "Pan"}
			precios := [3]float64{1.25, 0.95, 0.50}
			fmt.Println("Productos disponibles", productos, precios)
			fmt.Printf("Que producto quieres añadir: ")
			fmt.Println("1. Arroz")
			fmt.Println("2. Leche")
			fmt.Println("3. Pan")
			fmt.Scan(&usrOpcionProducto)
			fmt.Printf("Que cantidad has vendido: ")
			fmt.Scan(&cantidadVendida)
		} else if usrOpcion == 2 {
			fmt.Printf("Que quieres hacer: ")
			fmt.Println("1. Calcular el mostrar y total recaudado")
			fmt.Println("2. Mostrar Estadistica")
		} else {
			break
		}
	}
}
