package controller

import (
	"onlineStore/internal/entities"
	"reflect"
	"sort"
	"testing"
)

func Test_convertOrders(t *testing.T) {
	type args struct {
		entOrders []entities.Order
	}
	tests := []struct {
		name       string
		args       args
		wantOrders []Order
	}{
		{
			name: "",
			args: args{entOrders: []entities.Order{
				{
					Number: 10,
					Products: []entities.Product{
						{
							Id:       1,
							Name:     "Ноутбук",
							Quantity: 2,
							Shelf: entities.Shelf{
								Id:   1,
								Name: "А",
							},
							OptionalShelving: nil,
						},
						{
							Id:       3,
							Name:     "Телефон",
							Quantity: 1,
							Shelf: entities.Shelf{
								Id:   2,
								Name: "Б",
							},
							OptionalShelving: []entities.Shelf{
								{
									Id:   3,
									Name: "В",
								},
								{
									Id:   5,
									Name: "З",
								},
							},
						},
						{
							Id:       6,
							Name:     "Микрофон",
							Quantity: 1,
							Shelf: entities.Shelf{
								Id:   4,
								Name: "Ж",
							},
							OptionalShelving: nil,
						},
					},
				},
				{
					Number: 11,
					Products: []entities.Product{
						{
							Id:       2,
							Name:     "Телевизор",
							Quantity: 3,
							Shelf: entities.Shelf{
								Id:   1,
								Name: "А",
							},
							OptionalShelving: nil,
						},
					},
				},
			}},
			wantOrders: []Order{
				{
					Number:        10,
					ProductName:   "Ноутбук",
					ProductId:     1,
					Quantity:      2,
					MainShelf:     "А",
					OptionalShelf: nil,
				},
				{
					Number:      10,
					ProductName: "Телефон",
					ProductId:   3,
					Quantity:    1,
					MainShelf:   "Б",
					OptionalShelf: []string{
						"В", "З",
					},
				},
				{
					Number:        10,
					ProductName:   "Микрофон",
					ProductId:     6,
					Quantity:      1,
					MainShelf:     "Ж",
					OptionalShelf: nil,
				},
				{
					Number:        11,
					ProductName:   "Телевизор",
					ProductId:     2,
					Quantity:      3,
					MainShelf:     "А",
					OptionalShelf: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOrders := convertOrders(tt.args.entOrders)

			sort.Slice(gotOrders, func(i, j int) bool {
				return (gotOrders[i].Number < gotOrders[j].Number) && (gotOrders[i].ProductId < gotOrders[j].ProductId)
			})

			sort.Slice(tt.wantOrders, func(i, j int) bool {
				return (tt.wantOrders[i].Number < tt.wantOrders[j].Number) && (tt.wantOrders[i].ProductId < tt.wantOrders[j].ProductId)
			})

			for key, val := range gotOrders {
				sort.Slice(val.OptionalShelf, func(i, j int) bool {
					return val.OptionalShelf[i] < val.OptionalShelf[j]
				})
				gotOrders[key] = val
			}

			if !reflect.DeepEqual(gotOrders, tt.wantOrders) {
				t.Errorf("convertOrders() = \n%#v\n, want \n%#v", gotOrders, tt.wantOrders)
			}
		})
	}
}

func Test_sortOrdersByShelving(t *testing.T) {
	type args struct {
		orders []Order
	}
	tests := []struct {
		name string
		args args
		want map[string][]Order
	}{
		{
			name: "",
			args: args{orders: []Order{
				{ProductId: 1, MainShelf: "A"},
				{ProductId: 2, MainShelf: "Б"},
				{ProductId: 3, MainShelf: "A"},
				{ProductId: 4, MainShelf: "С"},
				{ProductId: 5, MainShelf: "С"},
			}},
			want: map[string][]Order{
				"A": {
					{ProductId: 1, MainShelf: "A"},
					{ProductId: 3, MainShelf: "A"},
				},
				"Б": {
					{ProductId: 2, MainShelf: "Б"},
				},
				"С": {
					{ProductId: 4, MainShelf: "С"},
					{ProductId: 5, MainShelf: "С"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortOrdersByShelving(tt.args.orders); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortOrdersByShelving() = %v, want %v", got, tt.want)
			}
		})
	}
}
