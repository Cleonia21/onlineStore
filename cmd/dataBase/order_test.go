package dataBase

import (
	"reflect"
	"testing"
)

func Test_dataBase_getShelving(t *testing.T) {
	db := DataBase{}
	db.ConnectDB()
	defer db.Close()

	//err := fillDB(&db)
	//if err != nil {
	//	fmt.Println(err)
	//}

	type args struct {
		ordersNum []int
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name:    "",
			args:    args{ordersNum: []int{12344}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr, _ := db.GetShelving(tt.args.ordersNum)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("getShelving() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
