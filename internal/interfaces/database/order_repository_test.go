package database

import (
	"onlineStore/internal/entities"
	"reflect"
	"testing"
)

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
			if gotKeys := keys(tt.args.m); !reflect.DeepEqual(gotKeys, tt.wantKeys) {
				t.Errorf("keys() = %v, want %v", gotKeys, tt.wantKeys)
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
