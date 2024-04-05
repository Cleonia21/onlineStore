package main

import (
	"fmt"
	"github.com/lib/pq"
)

type order struct {
	number        int
	productName   string
	productId     int
	quantity      int
	mainShelf     string
	optionalShelf []string
}

func (db *dataBase) getShelving(ordersNum []int) (err error, orders []order) {
	q := `
		SELECT
			o.number,
			p.name,
			p.id,
			o.quantity,
			s1.name AS "mainShelf",
			array_remove(array_agg(s2.name), NULL) AS "optionalShelf"
		FROM orders o
				 JOIN products p ON o.productId = p.id
				 JOIN shelving s1 ON p.shelfId = s1.id
				 LEFT JOIN optionalShelving os ON p.id = os.productId
				 LEFT JOIN shelving s2 ON os.shelfId = s2.id
		WHERE o.number IN %v
		GROUP BY o.number, o.quantity, p.name, p.id, s1.name
		`
	q = fmt.Sprintf(q, "("+intToStrJoin(ordersNum, ", ")+")")

	rows, err := db.db.Query(q)
	if err != nil {
		return fmt.Errorf("getShelving() %v", err.Error()), nil
	}
	defer rows.Close()

	for rows.Next() {
		o := order{}
		err = rows.Scan(
			&o.number,
			&o.productName,
			&o.productId,
			&o.quantity,
			&o.mainShelf,
			pq.Array(&o.optionalShelf),
		)
		if err != nil {
			return
		}
		orders = append(orders, o)
	}
	return
}
