package dataBase

import (
	"fmt"
	"math/rand"
	"time"
)

// cleanDb удаляет записи из всех таблиц и
// выставляет в них щетчик последовательности записей в дефолтное значение.
func cleanDb(db *DataBase) error {
	q := `
	ALTER SEQUENCE orders_id_seq RESTART WITH 1;
	delete from orders *;
	ALTER SEQUENCE optionalShelving_id_seq RESTART WITH 1;
	delete from optionalShelving *;
	ALTER SEQUENCE products_id_seq RESTART WITH 1;
	delete from products *;
	ALTER SEQUENCE shelving_id_seq RESTART WITH 1;
	delete from shelving *;
	`
	_, err := db.db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}

// fillDb заполняет таблицы в базе данных случайными значениями.
// Количество записей в таблицах <= entriesNum
func fillDb(db *DataBase, entriesNum int) error {
	rand.Seed(time.Now().UnixNano())

	if err := cleanDb(db); err != nil {
		return fmt.Errorf("cleanDb(): %v\n", err.Error())
	}
	if err := fillShelving(db, entriesNum); err != nil {
		return fmt.Errorf("fillShelving(): %v\n", err.Error())
	}
	fmt.Println("fillShelving() OK!")
	if err := fillProducts(db, entriesNum); err != nil {
		return fmt.Errorf("fillProducts(): %v\n", err.Error())
	}
	fmt.Println("fillProducts() OK!")
	if err := fillOptionalShelving(db, entriesNum); err != nil {
		return fmt.Errorf("fillOptionalShelving(): %v\n", err.Error())
	}
	fmt.Println("fillOptionalShelving() OK!")
	if err := fillOrders(db, entriesNum); err != nil {
		return fmt.Errorf("fillOrders(): %v\n", err.Error())
	}
	fmt.Println("fillOrders() OK!")
	return nil
}

func fillShelving(db *DataBase, entriesNum int) error {
	for i := 0; i < entriesNum; i++ {
		q := fmt.Sprintf("INSERT INTO shelving (name) VALUES ('%c')", randomLetter())
		_, err := db.db.Exec(q)
		if err != nil {
			return err
		}
	}
	return nil
}

func fillProducts(db *DataBase, entriesNum int) error {
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

func fillOptionalShelving(db *DataBase, entriesNum int) error {
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

func fillOrders(db *DataBase, entriesNum int) error {
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

// randomLetter генерирует случайную букву латинского алфавита
func randomLetter() rune {
	randomIndex := rand.Intn(26)
	return rune('a' + randomIndex)
}
