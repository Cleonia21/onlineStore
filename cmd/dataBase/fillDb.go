package dataBase

import (
	"fmt"
	"math/rand"
	"time"
)

const entriesNum = 1000000

func fillDB(db *DataBase) error {
	rand.Seed(time.Now().UnixNano())
	if err := fillShelving(db); err != nil {
		return fmt.Errorf("fillShelving() err=%v\n", err.Error())
	} else {
		fmt.Println("fillShelving() OK!")
	}
	if err := fillProducts(db); err != nil {
		return fmt.Errorf("fillProducts() err=%v\n", err.Error())
	}
	fmt.Println("fillProducts() OK!")
	if err := fillOptionalShelving(db); err != nil {
		return fmt.Errorf("fillOptionalShelving() err=%v\n", err.Error())
	}
	fmt.Println("fillOptionalShelving() OK!")
	if err := fillOrders(db); err != nil {
		return fmt.Errorf("fillOrders() err=%v\n", err.Error())
	}
	fmt.Println("fillOrders() OK!")
	return nil
}

func fillShelving(db *DataBase) error {
	for i := 0; i < entriesNum; i++ {
		q := fmt.Sprintf("INSERT INTO shelving (name) VALUES ('%c')", randomLetter())
		_, err := db.db.Exec(q)
		if err != nil {
			return err
		}
	}
	return nil
}

func fillProducts(db *DataBase) error {
	for i := 0; i < entriesNum; i++ {
		q := fmt.Sprintf(
			"INSERT INTO products (name, shelfId) VALUES ('%c', %v)",
			randomLetter(),
			rand.Intn(entriesNum-2)+1)
		_, err := db.db.Exec(q)
		if err != nil {
			return err
		}
	}
	return nil
}

func fillOptionalShelving(db *DataBase) error {
	for i := 0; i < entriesNum/10; i++ {
		q := fmt.Sprintf(
			"INSERT INTO optionalShelving (ProductId, shelfId) VALUES (%v, %v)",
			rand.Intn(entriesNum-2)+1,
			rand.Intn(entriesNum-2)+1)
		_, err := db.db.Exec(q)
		if err != nil {
			return err
		}
	}
	return nil
}

func fillOrders(db *DataBase) error {
	for i := 0; i < entriesNum/5; i++ {
		for j := 0; j < rand.Intn(5); j++ {
			q := fmt.Sprintf(
				"INSERT INTO orders (Number, ProductId, Quantity) VALUES (%v, %v, %v)",
				i,
				rand.Intn(entriesNum-2)+1,
				rand.Intn(10)+1)
			_, err := db.db.Exec(q)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func randomLetter() rune {
	// Генерируем случайное число от 0 до 25 (для букв a-z)
	randomIndex := rand.Intn(26)
	// Преобразуем индекс в букву, начиная с 'a'
	return rune('a' + randomIndex)
}
