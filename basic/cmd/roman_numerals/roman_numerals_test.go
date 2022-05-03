package roman_numerals

import "testing"

func TestRomanNumerals(t *testing.T) {
	tests := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
		{"4 gets converted to IV", 4, "IV"},
		{"5 gets converted to V", 5, "V"},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			got := ConvertRoman(tt.Arabic)
			if got != tt.Want {
				t.Errorf("👺 got %q, but want %q", got, tt.Want)
			}
		})
	}
}
