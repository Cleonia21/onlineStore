package infrastructure

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"onlineStore/internal/interfaces/database"
	"onlineStore/utils"
)

type SqlHandler struct {
	db *sql.DB
}

func NewSqlHandler() *SqlHandler {
	connStr := fmt.Sprintf("postgres://%v:%v@%v:5432/%v?sslmode=disable",
		"postgres",
		"2222",
		"localhost",
		"onlineStore",
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}

func (handler *SqlHandler) GetOrders(id int) ([]database.Order, error) {
	q := `SELECT productId, quantity FROM orders WHERE number = %v`
	q = fmt.Sprintf(q, id)

	rows, err := handler.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []database.Order
	for rows.Next() {
		var order database.Order
		err = rows.Scan(&order.ProductId, &order.Quantity)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, err
}

func (handler *SqlHandler) GetProducts(id []int) ([]database.Product, error) {
	q := `SELECT id, name, shelfId FROM products WHERE id IN (%v)`
	q = fmt.Sprintf(q, utils.IntToStrJoin(id, ","))

	rows, err := handler.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []database.Product
	for rows.Next() {
		var product database.Product
		err = rows.Scan(&product.Id, &product.Name, &product.ShelfId)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, err
}

func (handler *SqlHandler) GetShelving(id []int) ([]database.Shelf, error) {
	q := `SELECT id, name FROM shelving WHERE id IN (%v)`
	q = fmt.Sprintf(q, utils.IntToStrJoin(id, ","))

	rows, err := handler.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shelving []database.Shelf
	for rows.Next() {
		var shelf database.Shelf
		err = rows.Scan(&shelf.Id, &shelf.Name)
		if err != nil {
			return nil, err
		}
		shelving = append(shelving, shelf)
	}
	return shelving, err
}

func (handler *SqlHandler) GetOptionalShelving(ProductId []int) ([]database.OptionalShelving, error) {
	q := `SELECT productId, shelfId FROM optionalShelving WHERE productId IN (%v)`
	q = fmt.Sprintf(q, utils.IntToStrJoin(ProductId, ","))

	rows, err := handler.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var optionalShelving []database.OptionalShelving
	for rows.Next() {
		var optionalShelf database.OptionalShelving
		err = rows.Scan(&optionalShelf.ProductId, &optionalShelf.ShelfId)
		if err != nil {
			return nil, err
		}
		optionalShelving = append(optionalShelving, optionalShelf)
	}
	return optionalShelving, err
}
