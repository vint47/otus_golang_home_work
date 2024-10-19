package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var res strings.Builder
	for i := 0; i < len(s); i++ {
		ch := string(s[i])
		countCh := 1

		if i+1 < len(s) {
			chNext := string(s[i+1])
			if numb, err := strconv.Atoi(chNext); err == nil {
				if _, err := strconv.Atoi(ch); err == nil {
					return "", ErrInvalidString
				}
				countCh = numb
			}
		}

		if _, err := strconv.Atoi(ch); err != nil {
			res.WriteString(strings.Repeat(ch, countCh))
		} else if i == 0 {
			return "", ErrInvalidString
		}
	}

	return res.String(), nil
}
