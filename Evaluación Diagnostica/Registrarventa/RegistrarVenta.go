package registrarventa

func RegistrarVenta(productos string, precios float64, usrOpcionProducto int, cantidadVendida float64) float64 {
	if usrOpcionProducto == "1" {
		precioarroz := precios[0]
		return float64(precioarroz) * cantidadVendida
		append(productos[1])
	} else if usrOpcionProducto == "2" {
		precioleche := precios[1]
		return float64(precioleche) * cantidadVendida
	} else {
		preciopan := precios[2]
		return float64(preciopan) * cantidadVendida
	}

}
