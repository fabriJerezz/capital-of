package main

import (
	"ci-environment/api"
	// "ci-environment/utils"
	"fmt"
	"log"
	"net/http"

)

func main() {
	http.HandleFunc("/obtener-capital", api.Handler)            // Define la ruta raiz
    fmt.Println("Servidor en http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)        // Escucha en el puerto 8080
    
    if err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
	// fmt.Print("Ingresa el nombre de una provincia de la que quieras saber su capital: ") 
	//// fmt.Scanln(&province)
}
