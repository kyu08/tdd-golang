package roman_numerals

import "testing"

func TestRomanNumerals(t *testing.T) {
	got := ConvertRoman(1)
	want := "I"

	if got != want {
		t.Errorf("👺 got %q, but want %q", got, want)
	}
}
