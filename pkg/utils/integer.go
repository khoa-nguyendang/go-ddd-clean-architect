package utils

import (
	"strconv"
)

func GetInteger(input string) int {
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}
	return number
}
