package utils

import (
	"testing"
)


func TestGetCapital (t *testing.T) {
	tests := []struct {
		name     string
		province     string
		expectedCapital string
	}{
		{"real province", "Chaco", "Resistencia"},
		{"case and accent marks", "sÁnTá fÉ", "Santa Fe"},
		{"another word", "Chiavenato", "No se encontró la provincia Chiavenato"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			capital, err := GetCapital(tt.province)
			
			if err != nil {
				t.Errorf("GetCapital(%s) = %s, %v; want %s, nil", tt.province, capital, err, tt.expectedCapital)
			}
			if capital != tt.expectedCapital {
				t.Errorf("GetCapital(%s) = %s; want %s", tt.province, capital, tt.expectedCapital)
			}
		})
	}
}