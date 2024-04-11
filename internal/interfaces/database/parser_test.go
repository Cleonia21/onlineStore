package database

import (
	"onlineStore/internal/entities"
	"reflect"
	"sort"
	"testing"
)

func Test_parser_saveOrders(t *testing.T) {
	type fields struct {
		productsMap map[int]entities.Product
		shelvingMap map[int]entities.Shelf
	}
	type args struct {
		orders []Order
	}
	tests := []struct {
		name       string
		args       args
		fields     fields
		wantFields fields
	}{
		{
			name: "",
			args: args{orders: []Order{
				{
					ProductId: 1,
					Quantity:  65331,
				},
				{
					ProductId: 2,
					Quantity:  1232,
				},
			}},
			fields: fields{
				productsMap: map[int]entities.Product{
					3: {
						Id:       3,
						Quantity: 321,
					},
				},
			},
			wantFields: fields{
				productsMap: map[int]entities.Product{
					1: {
						Id:       1,
						Quantity: 65331,
					},
					2: {
						Id:       2,
						Quantity: 1232,
					},
					3: {
						Id:       3,
						Quantity: 321,
					},
				},
			},
		},
		{
			name: "",
			args: args{orders: nil},
			fields: fields{
				productsMap: map[int]entities.Product{
					3: {
						Id:       3,
						Quantity: 321,
					},
				},
			},
			wantFields: fields{
				productsMap: map[int]entities.Product{
					3: {
						Id:       3,
						Quantity: 321,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &parser{
				productsMap: tt.fields.productsMap,
				shelvingMap: tt.fields.shelvingMap,
			}
			s.saveOrders(tt.args.orders)

			gotFields := fields{
				productsMap: s.productsMap,
				shelvingMap: s.shelvingMap,
			}
			if !reflect.DeepEqual(tt.fields, tt.wantFields) {
				t.Errorf("gotFields = %v, wantFields %v", gotFields, tt.wantFields)
			}
		})
	}
}

func Test_parser_saveProducts(t *testing.T) {
	type fields struct {
		productsMap map[int]entities.Product
		shelvingMap map[int]entities.Shelf
	}
	type args struct {
		products []Product
	}
	tests := []struct {
		name       string
		args       args
		fields     fields
		wantFields fields
	}{
		{
			name: "",
			args: args{products: []Product{
				{
					Id:      1,
					Name:    "qwe",
					ShelfId: 1,
				},
				{
					Id:      2,
					Name:    "rty",
					ShelfId: 3,
				},
			}},
			fields: fields{
				productsMap: map[int]entities.Product{},
				shelvingMap: map[int]entities.Shelf{},
			},
			wantFields: fields{
				productsMap: map[int]entities.Product{
					1: {
						Name:  "qwe",
						Shelf: entities.Shelf{Id: 1},
					},
					2: {
						Name:  "rty",
						Shelf: entities.Shelf{Id: 3},
					},
				},
				shelvingMap: map[int]entities.Shelf{
					3: entities.Shelf{Id: 3},
					1: entities.Shelf{Id: 1},
				},
			},
		},
		{
			name: "",
			args: args{products: nil},
			fields: fields{
				productsMap: map[int]entities.Product{
					3: {
						Id:       3,
						Quantity: 321,
					},
				},
			},
			wantFields: fields{
				productsMap: map[int]entities.Product{
					3: {
						Id:       3,
						Quantity: 321,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &parser{
				productsMap: tt.fields.productsMap,
				shelvingMap: tt.fields.shelvingMap,
			}
			s.saveProducts(tt.args.products)

			gotFields := fields{
				productsMap: s.productsMap,
				shelvingMap: s.shelvingMap,
			}
			if !reflect.DeepEqual(tt.fields.productsMap, tt.wantFields.productsMap) {
				t.Errorf("gotProductsMap = %v, wantProductsMap %v", gotFields.productsMap, tt.wantFields.productsMap)
			}
			if !reflect.DeepEqual(tt.fields.shelvingMap, tt.wantFields.shelvingMap) {
				t.Errorf("gotShelvingMap= %v, wantShelvingMap %v", gotFields.shelvingMap, tt.wantFields.shelvingMap)
			}
		})
	}
}

func Test_parser_saveOptShelving(t *testing.T) {
	type fields struct {
		productsMap map[int]entities.Product
		shelvingMap map[int]entities.Shelf
	}
	type args struct {
		optionalShelving []OptionalShelving
	}
	tests := []struct {
		name       string
		args       args
		fields     fields
		wantFields fields
	}{
		{
			name: "",
			args: args{optionalShelving: []OptionalShelving{
				{
					ProductId: 1,
					ShelfId:   1,
				},
				{
					ProductId: 1,
					ShelfId:   2,
				},
				{
					ProductId: 1,
					ShelfId:   3,
				},
				{
					ProductId: 12,
					ShelfId:   1,
				},
				{
					ProductId: 14,
					ShelfId:   1,
				},
			}},
			fields: fields{
				productsMap: map[int]entities.Product{
					1:  {},
					12: {},
					14: {},
				},
				shelvingMap: map[int]entities.Shelf{
					1: {Id: 1},
				},
			},
			wantFields: fields{
				productsMap: map[int]entities.Product{
					1: {
						OptionalShelving: []entities.Shelf{
							{Id: 1},
							{Id: 2},
							{Id: 3},
						},
					},
					12: {
						OptionalShelving: []entities.Shelf{
							{Id: 1},
						},
					},
					14: {
						OptionalShelving: []entities.Shelf{
							{Id: 1},
						},
					},
				},
				shelvingMap: map[int]entities.Shelf{
					1: {Id: 1},
					2: {Id: 2},
					3: {Id: 3},
				},
			},
		},
		{
			name: "",
			args: args{optionalShelving: nil},
			fields: fields{
				productsMap: map[int]entities.Product{
					3: {
						Id:       3,
						Quantity: 321,
					},
				},
			},
			wantFields: fields{
				productsMap: map[int]entities.Product{
					3: {
						Id:       3,
						Quantity: 321,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &parser{
				productsMap: tt.fields.productsMap,
				shelvingMap: tt.fields.shelvingMap,
			}
			s.saveOptionalShelving(tt.args.optionalShelving)

			gotFields := fields{
				productsMap: s.productsMap,
				shelvingMap: s.shelvingMap,
			}
			if !reflect.DeepEqual(tt.fields.productsMap, tt.wantFields.productsMap) {
				t.Errorf("gotProductsMap = %v, wantProductsMap %v", gotFields.productsMap, tt.wantFields.productsMap)
			}
			if !reflect.DeepEqual(tt.fields.shelvingMap, tt.wantFields.shelvingMap) {
				t.Errorf("gotShelvingMap= %v, wantShelvingMap %v", gotFields.shelvingMap, tt.wantFields.shelvingMap)
			}
		})
	}
}

func Test_parser_saveShelving(t *testing.T) {
	type args struct {
		shelving []Shelf
	}
	type want struct {
		m map[int]entities.Shelf
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "",
			args: args{shelving: []Shelf{
				{
					Id:   1,
					Name: "1",
				},
				{
					Id:   123,
					Name: "123",
				},
			}},
			want: want{m: map[int]entities.Shelf{
				1: {
					Id:   1,
					Name: "1",
				},
				123: {
					Id:   123,
					Name: "123",
				},
			}},
		},
		{
			name: "",
			args: args{nil},
			want: want{m: map[int]entities.Shelf{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := initParser()
			s.saveShelving(tt.args.shelving)
			if !reflect.DeepEqual(s.shelvingMap, tt.want.m) {
				t.Errorf("shelvingMap() = %v, want %v", s.shelvingMap, tt.want.m)
			}
		})
	}
}

func Test_parser_parseData(t *testing.T) {
	type fields struct {
		productsMap map[int]entities.Product
		shelvingMap map[int]entities.Shelf
	}
	tests := []struct {
		name         string
		fields       fields
		wantProducts []entities.Product
	}{
		{
			name: "",
			fields: fields{
				productsMap: map[int]entities.Product{
					1: {
						Id:       1,
						Name:     "Компьютер",
						Quantity: 2,
						Shelf: entities.Shelf{
							Id:   1,
							Name: "",
						},
						OptionalShelving: []entities.Shelf{
							{
								Id:   2,
								Name: "",
							},
							{
								Id:   3,
								Name: "",
							},
							{
								Id:   4,
								Name: "",
							},
						},
					},
					12: {
						Id:       12,
						Name:     "Телефон",
						Quantity: 21,
						Shelf: entities.Shelf{
							Id:   4,
							Name: "",
						},
						OptionalShelving: nil,
					},
				},
				shelvingMap: map[int]entities.Shelf{
					1: {
						Id:   1,
						Name: "A",
					},
					2: {
						Id:   2,
						Name: "B",
					},
					3: {
						Id:   3,
						Name: "C",
					},
					4: {
						Id:   4,
						Name: "O",
					},
				},
			},
			wantProducts: []entities.Product{
				{
					Id:       1,
					Name:     "Компьютер",
					Quantity: 2,
					Shelf: entities.Shelf{
						Id:   1,
						Name: "A",
					},
					OptionalShelving: []entities.Shelf{
						{
							Id:   2,
							Name: "B",
						},
						{
							Id:   3,
							Name: "C",
						},
						{
							Id:   4,
							Name: "O",
						},
					},
				},
				{
					Id:       12,
					Name:     "Телефон",
					Quantity: 21,
					Shelf: entities.Shelf{
						Id:   4,
						Name: "O",
					},
					OptionalShelving: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &parser{
				productsMap: tt.fields.productsMap,
				shelvingMap: tt.fields.shelvingMap,
			}
			gotProducts := s.parseData()

			sort.Slice(gotProducts, func(i, j int) bool {
				return gotProducts[i].Id < gotProducts[j].Id
			})

			if !reflect.DeepEqual(gotProducts, tt.wantProducts) {
				t.Errorf("parseData() = \n%#v\n, want \n%#v", gotProducts, tt.wantProducts)
			}
		})
	}
}

func Test_keys(t *testing.T) {
	type args[T any] struct {
		m map[int]T
	}
	type testCase[T any] struct {
		name     string
		args     args[T]
		wantKeys []int
	}
	tests := []testCase[int]{
		{
			name: "",
			args: args[int]{
				m: map[int]int{
					1: 23423,
					2: 2432,
					3: 14213,
				},
			},
			wantKeys: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKeys := keys(tt.args.m)

			sort.Slice(gotKeys, func(i, j int) bool {
				return gotKeys[i] < gotKeys[j]
			})

			if !reflect.DeepEqual(gotKeys, tt.wantKeys) {
				t.Errorf("keys() = %v, want %v", gotKeys, tt.wantKeys)
			}
		})
	}
}
