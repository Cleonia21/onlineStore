package utils

import (
	"strconv"
	"strings"
)

// IntToStrJoin работает как strings.Join(), но для массива чисел
func IntToStrJoin(elems []int, sep string) string {
	var strs []string
	for _, i := range elems {
		strs = append(strs, strconv.Itoa(i))
	}
	return strings.Join(strs, sep)
}
