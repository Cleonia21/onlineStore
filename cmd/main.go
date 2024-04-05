package main

import (
	"fmt"
	"main/dataBase"
)

func main() {
	ordersNum := []int{10, 11, 14, 15}

	db := dataBase.DataBase{}
	db.ConnectDB()
	defer db.Close()

	err, orders := db.GetShelving(ordersNum)
	if err != nil {
		fmt.Println(err)
		return
	}

	myPrint(ordersNum, sortByShelving(orders))
}

func sortByShelving(orders []dataBase.Order) map[string][]dataBase.Order {
	sortMap := make(map[string][]dataBase.Order)
	for _, o := range orders {
		sortMap[o.MainShelf] = append(sortMap[o.MainShelf], o)
	}
	return sortMap
}
