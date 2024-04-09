package main

import (
	"fmt"
	"onlineStore/dataBase"
	"onlineStore/utils"
	"strings"
)

func printOrders(ordersNum []int, mapByShelving map[string][]dataBase.Order) {
	fmt.Printf(
		"=+=+=+=\nСтраница сборки заказов %v\n",
		utils.IntToStrJoin(ordersNum, ","),
	)
	for key, val := range mapByShelving {
		printShelf(key, val)
	}
}

func printShelf(shelf string, orders []dataBase.Order) {
	fmt.Printf("===Стеллаж %v\n", shelf)
	for _, o := range orders {
		printProduct(o)
	}
}

func printProduct(o dataBase.Order) {
	fmt.Printf("%v (id=%v)\n", o.ProductName, o.ProductId)
	fmt.Printf("заказ %v, %v шт", o.Number, o.Quantity)
	if len(o.OptionalShelf) != 0 {
		fmt.Printf("\nдоп стеллаж: %v", strings.Join(o.OptionalShelf, ","))
	}
	fmt.Printf("\n\n")
}
