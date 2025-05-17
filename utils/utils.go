package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func FindCapitalCity(province string) string {
	jsonContent, err := os.ReadFile("capital_cities.json")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	var capitalCities map[string]string
	err = json.Unmarshal(jsonContent, &capitalCities)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return ""
	}
	capitalCity := capitalCities[province]
	return capitalCity
}
