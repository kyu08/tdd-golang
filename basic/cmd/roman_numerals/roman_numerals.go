package roman_numerals

import "strings"

func ConvertRoman(arabic int) string {
	var result strings.Builder

	for _, n := range allRomanNumerals {
		for arabic >= n.value {
			result.WriteString(n.symbol)
			arabic -= n.value
		}
	}

	return result.String()
}

type RomanNumeral struct {
	value  int
	symbol string
}

var allRomanNumerals = []RomanNumeral{
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}
