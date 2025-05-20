package utils

import (
	"testing"
)


func TestGetCapital () {
	tests := []struct {
		name     string
		province     string
		expectedCapital string
	}{
		{"real province", "Chaco", "Resistencia"},
		{"case and accent marks", "sÁnTá fÉ", "Santa Fé"},
		{"another word", "Chiavenato", "No se encontró la provincia Chiavenat"},
	}
}