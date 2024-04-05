package main

import (
	"strconv"
	"strings"
)

func intToStrJoin(ordersNum []int, sep string) string {
	var strs []string
	for _, i := range ordersNum {
		strs = append(strs, strconv.Itoa(i))
	}
	return strings.Join(strs, sep)
}
