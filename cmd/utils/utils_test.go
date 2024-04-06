package utils

import "testing"

func TestIntToStrJoin(t *testing.T) {
	type args struct {
		elems []int
		sep   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				elems: []int{1, 2, 3},
				sep:   ",",
			},
			want: "1,2,3",
		},
		{
			name: "",
			args: args{
				elems: nil,
				sep:   ",",
			},
			want: "",
		},
		{
			name: "",
			args: args{
				elems: []int{1},
				sep:   ",",
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToStrJoin(tt.args.elems, tt.args.sep); got != tt.want {
				t.Errorf("IntToStrJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}
