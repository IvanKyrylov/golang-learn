package tasktwo

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// UnpackingString

var ErrInvalidString = errors.New("invalid string")

func UnpackingString(str string) (string, error) {

	if len(str) == 0 {
		return "", nil
	}

	tempRuneStr := []rune(str)
	var builder strings.Builder

	if unicode.IsDigit(tempRuneStr[0]) {
		return "", ErrInvalidString
	}

	for i, v := range tempRuneStr {

		if unicode.IsDigit(v) {
			if unicode.IsDigit(tempRuneStr[i-1]) {
				return "", ErrInvalidString
			}
			count, _ := strconv.Atoi(string(v))
			if count > 1 {
				builder.WriteString(strings.Repeat(string(tempRuneStr[i-1]), count-1))
			}
			if count == 0 {
				s := builder.String()[:builder.Len()-1]
				builder.Reset()
				builder.WriteString(s)
			}
		} else {
			builder.WriteRune(v)
		}
	}

	return builder.String(), nil
}
