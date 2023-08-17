package drae

import (
	"errors"
	"strconv"
	"strings"
)

func SanitizeStrNum(strNum string) (int, error) {
	strNum = strings.TrimFunc(strNum, func(r rune) bool {
		if r == ' ' || r == '.' {
			return true
		}
		return false
	})

	num, err := strconv.Atoi(strNum)
	if err != nil {
		return 0, errors.New("invalid number")
	}

	return num, nil
}
