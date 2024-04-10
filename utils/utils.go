package utils

import (
	"strconv"
	"strings"
)

func IntToStrJoin(ordersNum []int, sep string) string {
	var strs []string
	for _, i := range ordersNum {
		strs = append(strs, strconv.Itoa(i))
	}
	return strings.Join(strs, sep)
}
