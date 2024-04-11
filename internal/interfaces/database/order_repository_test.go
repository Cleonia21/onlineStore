package database

import (
	"onlineStore/internal/entities"
	"reflect"
	"sort"
	"testing"
)

type repo1 struct{}

func (r repo1) GetOrders(id int) ([]Order, error) {
	_ = id
	return []Order{
		{
			ProductId: 1,
			Quantity:  2,
		},
		{
			ProductId: 3,
			Quantity:  1,
		},
		{
			ProductId: 6,
			Quantity:  1,
		},
	}, nil
}

func (r repo1) GetProducts(id []int) ([]Product, error) {
	_ = id
	return []Product{
		{
			Id:      1,
			Name:    "Ноутбук",
			ShelfId: 1, // А
		},
		{
			Id:      3,
			Name:    "Телефон",
			ShelfId: 2, // Б
		},
		{
			Id:      6,
			Name:    "Микрофон",
			ShelfId: 4, // Ж
		},
	}, nil
}

func (r repo1) GetShelving(id []int) ([]Shelf, error) {
	_ = id
	return []Shelf{
		{
			Id:   1,
			Name: "А",
		},
		{
			Id:   2,
			Name: "Б",
		},
		{
			Id:   3,
			Name: "В",
		},
		{
			Id:   4,
			Name: "Ж",
		},
		{
			Id:   5,
			Name: "З",
		},
	}, nil
}

func (r repo1) GetOptionalShelving(ProductId []int) ([]OptionalShelving, error) {
	_ = ProductId
	return []OptionalShelving{
		{
			ProductId: 3,
			ShelfId:   3,
		},
		{
			ProductId: 3,
			ShelfId:   5,
		},
	}, nil
}

func TestOrderRepository_SelectOrder(t *testing.T) {
	type fields struct {
		SqlHandler SqlHandler
	}
	type args struct {
		number int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.Order
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				SqlHandler: repo1{},
			},
			args: args{number: 1},
			want: entities.Order{
				Number: 1,
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			or := &OrderRepository{
				SqlHandler: tt.fields.SqlHandler,
			}
			got, err := or.SelectOrder(tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			sort.Slice(got.Products, func(i, j int) bool {
				return got.Products[i].Id < got.Products[j].Id
			})

			for key, val := range got.Products {
				sort.Slice(val.OptionalShelving, func(i, j int) bool {
					return val.OptionalShelving[i].Id < val.OptionalShelving[j].Id
				})
				got.Products[key] = val
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectOrder() got = \n%#v\n, want \n%#v", got, tt.want)
			}
		})
	}
}
