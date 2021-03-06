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

func ConvertToArabic(roman string) int {
	arabic := 0
	for range roman {
		arabic++
	}
	return arabic
}

type romanNumeral struct {
	value  int
	symbol string
}

var allRomanNumerals = []romanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}
