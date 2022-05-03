package roman_numerals

import "strings"

func ConvertRoman(arabic int) string {

	var result strings.Builder
	for i := arabic; i > 0; i-- {
		if arabic == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
	}
	return result.String()
}
