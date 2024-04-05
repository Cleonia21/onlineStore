package main

import (
	"fmt"
	"strings"
)

func myPrint(ordersNum []int, mapByShelving map[string][]order) {
	fmt.Printf(
		"=+=+=+=\nСтраница сборки заказов %v\n",
		intToStrJoin(ordersNum, ","),
	)
	for key, val := range mapByShelving {
		printShelf(key, val)
	}
}

func printShelf(shelf string, orders []order) {
	fmt.Printf("===Стеллаж %v\n", shelf)
	for _, o := range orders {
		printProduct(o)
	}
}

func printProduct(o order) {
	fmt.Printf("%v (id=%v)\n", o.productName, o.productId)
	fmt.Printf("заказ %v, %v шт", o.number, o.quantity)
	if len(o.optionalShelf) != 0 {
		fmt.Printf("\nдоп стеллаж: %v", strings.Join(o.optionalShelf, ","))
	}
	fmt.Printf("\n\n")
}
