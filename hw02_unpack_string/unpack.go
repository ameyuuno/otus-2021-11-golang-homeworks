package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	unpackedString := strings.Builder{}

	var previousSymbol rune
	for _, currentSymbol := range str {
		if previousSymbol == 0 {
			if unicode.IsDigit(currentSymbol) {
				return "", ErrInvalidString
			}
			previousSymbol = currentSymbol
			continue
		}

		if !unicode.IsDigit(currentSymbol) {
			unpackedString.WriteRune(previousSymbol)
			previousSymbol = currentSymbol
			continue
		}

		repetitionNumber, err := strconv.Atoi(string(currentSymbol))
		if err != nil {
			return "", ErrInvalidString
		}

		unpackedString.WriteString(strings.Repeat(string(previousSymbol), repetitionNumber))
		previousSymbol = 0
	}

	if previousSymbol != 0 {
		unpackedString.WriteRune(previousSymbol)
	}

	return unpackedString.String(), nil
}
