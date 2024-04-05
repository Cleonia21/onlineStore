package main

import (
	"fmt"
)

func main() {
	ordersNum := []int{10, 11, 14, 15}

	db := dataBase{}
	db.ConnectDB()
	defer db.Close()

	err, orders := db.getShelving(ordersNum)
	if err != nil {
		fmt.Println(err)
		return
	}

	myPrint(ordersNum, sortByShelving(orders))
}

func sortByShelving(orders []order) map[string][]order {
	sortMap := make(map[string][]order)
	for _, o := range orders {
		sortMap[o.mainShelf] = append(sortMap[o.mainShelf], o)
	}
	return sortMap
}
