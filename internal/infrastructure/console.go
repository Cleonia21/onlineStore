package infrastructure

import (
	"errors"
	"fmt"
	"log"
	"onlineStore/internal/interfaces/controller"
	"onlineStore/utils"
	"os"
	"strconv"
	"strings"
)

func RunConsole() {
	ordersId, err := readArgv()
	if err != nil {
		log.Println(err)
		return
	}

	orderController := controller.NewOrderController(NewSqlHandler())

	orders, err := orderController.GetOrders(ordersId)
	if err != nil {
		log.Println(err)
		return
	}

	printOrders(ordersId, orders)
}

func readArgv() (ordersId []int, err error) {
	err = errors.New("не верные аргументы при запуске программы (go run main.go 10,11,14,15)")
	var intArr []int
	var strArr []string

	if len(os.Args) != 2 {
		return
	}

	strArr = strings.Split(os.Args[1], ",")
	for _, arg := range strArr {
		n, err := strconv.Atoi(arg)
		if err == nil && n > 0 {
			intArr = append(intArr, n)
		}
	}
	if len(intArr) < 1 {
		return
	}
	return intArr, nil
}

func printOrders(ordersNum []int, mapByShelving map[string][]controller.Order) {
	fmt.Printf(
		"=+=+=+=\nСтраница сборки заказов %v\n",
		utils.IntToStrJoin(ordersNum, ","),
	)
	for key, val := range mapByShelving {
		printShelf(key, val)
	}
}

func printShelf(shelf string, orders []controller.Order) {
	fmt.Printf("===Стеллаж %v\n", shelf)
	for _, o := range orders {
		printProduct(o)
	}
}

func printProduct(o controller.Order) {
	fmt.Printf("%v (id=%v)\n", o.ProductName, o.ProductId)
	fmt.Printf("заказ %v, %v шт", o.Number, o.Quantity)
	if len(o.OptionalShelf) != 0 {
		fmt.Printf("\nдоп стеллаж: %v", strings.Join(o.OptionalShelf, ","))
	}
	fmt.Printf("\n\n")
}
