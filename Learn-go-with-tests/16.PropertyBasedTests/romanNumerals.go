package romanNumerals

import (
	"strings"
)


type RomanNumeral struct {
	Value uint16
	Symbol string
}

var allRomanNumerals = RomanNumerals {
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

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

func ConvertToRoman(number uint16) string {
	var result strings.Builder

	for _, romanNumeral := range allRomanNumerals {
		for number >= romanNumeral.Value {
			result.WriteString(romanNumeral.Symbol)
			number -= romanNumeral.Value
		}
	}

	return result.String()
}

var ValidSubtractors = []string{"I", "X", "C"}

func IsValidSubtractor(symbol string) bool {
	for _, subtractor := range ValidSubtractors {
		if symbol == subtractor {
			return true
		}
	}
	return false
}

func ConvertToArabic(roman string) uint16 {
	var total uint16

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if i+1 < len(roman) && IsValidSubtractor(string(symbol)) && allRomanNumerals.ValueOf(symbol) < allRomanNumerals.ValueOf(roman[i+1]) {

			value := allRomanNumerals.ValueOf(roman[i], roman[i+1])

			if value != 0 {
				total += value
				i++
			} else {
				total++
			}
		} else {
			total += allRomanNumerals.ValueOf(symbol)
		}
	}

	return total
}