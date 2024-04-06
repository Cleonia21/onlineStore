package dataBase

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_dataBase_getShelving(t *testing.T) {
	db := DataBase{}
	err := db.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	type args struct {
		ordersNum []int
	}
	tests := []struct {
		name     string
		args     args
		wantErr  error
		reFillDb bool
	}{
		{
			name:     "",
			args:     args{ordersNum: []int{10, 11, 14, 15}},
			wantErr:  nil,
			reFillDb: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.reFillDb {
				err = fillDb(&db, 1000)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			_, gotErr := db.GetOrders(tt.args.ordersNum)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("getShelving() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
