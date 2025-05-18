package main

import (
	"ci-environment/api"
	"ci-environment/utils"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", api.Handler)            // Define la ruta raíz
    fmt.Println("Servidor corriendo en http://localhost:8080")
    http.ListenAndServe(":8080", nil)        // Escucha en el puerto 8080
	
	var province string
	// fmt.Print("Ingresa el nombre de una provincia de la que quieras saber su capital: ")
	// fmt.Scanln(&province)

	capitalCity, err := utils.FindCapitalCity(province)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("La capital es", capitalCity)
}
