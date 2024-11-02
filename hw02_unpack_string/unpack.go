package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var res strings.Builder
	sl := []rune(s)

	if len(sl) > 0 {
		if _, err := strconv.Atoi(string(sl[0])); err == nil {
			return "", ErrInvalidString
		}
	}

	for i, v := range sl {
		ch := string(v)
		countCh := 1
		_, characterNumbErr := strconv.Atoi(ch)

		if i+1 < len(sl) {
			chNext := string(sl[i+1])
			if numb, err := strconv.Atoi(chNext); err == nil {
				if characterNumbErr == nil {
					return "", ErrInvalidString
				}
				countCh = numb
			}
		}

		if characterNumbErr != nil {
			res.WriteString(strings.Repeat(ch, countCh))
		}
	}

	return res.String(), nil
}
