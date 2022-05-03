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
		{"6 gets converted to VI", 6, "VI"},
		{"9 gets converted to IX", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
		{"14 gets converted to XIV", 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
		{"40 gets converted to XL", 40, "XL"},
		{"47 gets converted to XLVII", 47, "XLVII"},
		{"49 gets converted to XLIX", 49, "XLIX"},
		{"50 gets converted to L", 50, "L"},
		{"90", 90, "XC"},
		{"100", 100, "C"},
		{"400", 400, "CD"},
		{"451", 451, "CDLI"},
		{"500", 500, "D"},
		{"900", 900, "CM"},
		{"1000", 1000, "M"},
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
