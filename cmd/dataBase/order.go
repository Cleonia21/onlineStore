package dataBase

import (
	"errors"
	"fmt"
	"github.com/lib/pq"
	"main/utils"
)

type Order struct {
	Number        int
	ProductName   string
	ProductId     int
	Quantity      int
	MainShelf     string
	OptionalShelf []string
}

func (db *DataBase) GetOrders(ordersNum []int) (orders []Order, err error) {
	q := `
		SELECT
			o.Number,
			p.name,
			p.id,
			o.Quantity,
			s1.name AS "MainShelf",
			array_remove(array_agg(s2.name), NULL) AS "OptionalShelf"
		FROM orders o
				 JOIN products p ON o.ProductId = p.id
				 JOIN shelving s1 ON p.shelfId = s1.id
				 LEFT JOIN optionalShelving os ON p.id = os.ProductId
				 LEFT JOIN shelving s2 ON os.shelfId = s2.id
		WHERE o.Number IN (%v)
		GROUP BY o.Number, o.Quantity, p.name, p.id, s1.name
		`
	q = fmt.Sprintf(q, utils.IntToStrJoin(ordersNum, ", "))

	rows, err := db.db.Query(q)
	if err != nil {
		return nil, fmt.Errorf("getShelving() %v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		o := Order{}
		err = rows.Scan(
			&o.Number,
			&o.ProductName,
			&o.ProductId,
			&o.Quantity,
			&o.MainShelf,
			pq.Array(&o.OptionalShelf),
		)
		if err != nil {
			return
		}
		orders = append(orders, o)
	}
	if len(orders) == 0 {
		return nil, errors.New("заказы не найдены")
	}
	return
}
