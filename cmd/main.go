package main

import (
	"errors"
	"fmt"
	"main/dataBase"
	"os"
	"strconv"
	"strings"
)

func main() {
	ordersNum, err := parseArgv()
	if err != nil {
		fmt.Println(err)
		return
	}

	db := dataBase.DataBase{}
	err = db.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	orders, err := db.GetOrders(ordersNum)
	if err != nil {
		fmt.Println(err)
		return
	}

	printOrders(ordersNum, sortByShelving(orders))
}

// parseArgv считывает аргумент командной строки и возвращает его в виде массива чисел
func parseArgv() ([]int, error) {
	err := errors.New("не верные аргументы при запуске программы (go run main.go 10,11,14,15)")
	var intArr []int
	var strArr []string

	if len(os.Args) != 2 {
		return nil, err
	}

	strArr = strings.Split(os.Args[1], ",")
	for _, arg := range strArr {
		n, err := strconv.Atoi(arg)
		if err == nil && n > 0 {
			intArr = append(intArr, n)
		}
	}
	if len(intArr) < 1 {
		return nil, err
	}
	return intArr, nil
}

// sortByShelving сортирует массив заказов по полкам и
// возвращает map, где ключ стеллаж, а значение - массив заказов с этого стелажа
func sortByShelving(orders []dataBase.Order) map[string][]dataBase.Order {
	sortMap := make(map[string][]dataBase.Order)
	for _, o := range orders {
		sortMap[o.MainShelf] = append(sortMap[o.MainShelf], o)
	}
	return sortMap
}
