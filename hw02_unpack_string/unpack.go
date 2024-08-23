package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var (
		result     strings.Builder
		lastSymbol rune
		isSkip     = false
	)

	for _, charCode := range str {
		switch {
		case isSkip:
			lastSymbol = charCode
			isSkip = false

		case charCode == '\\':
			isSkip = true
			if lastSymbol != 0 {
				result.WriteRune(lastSymbol)
			}

		case unicode.IsDigit(charCode):
			count, err := strconv.Atoi(string(charCode))
			if err != nil {
				return "", err
			}
			if lastSymbol == 0 {
				return "", ErrInvalidString
			}
			if count > 0 {
				result.WriteString(strings.Repeat(string(lastSymbol), count))
				lastSymbol = 0
			} else {
				lastSymbol = 0
			}
		default:
			if lastSymbol != 0 {
				result.WriteRune(lastSymbol)
			}
			lastSymbol = charCode
		}
	}
	if lastSymbol != 0 {
		result.WriteRune(lastSymbol)
	}
	return result.String(), nil
}
