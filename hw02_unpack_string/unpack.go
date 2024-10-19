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

	for i, v := range sl {
		ch := string(v)
		countCh := 1

		if i+1 < len(sl) {
			chNext := string(sl[i+1])
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
