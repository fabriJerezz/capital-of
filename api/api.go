package api

import (
	"net/http"
	"fmt"
	"ci-environment/utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		province := r.URL.Query().Get("provincia")
		capital, err := utils.GetCapital(province)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Controlar si no ponen el parámetro
		// Controlar si ponen otro parámetro
		fmt.Fprint(w, "La capital es ", capital)
	} else {
		fmt.Fprintf(w, "Método http no válido")
	}
}