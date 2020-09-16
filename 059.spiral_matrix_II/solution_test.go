package leetcode

import (
	"reflect"
	"testing"
)

func Test_generateMatrix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "t5",
			args: args{
				n: 5,
			},
			want: [][]int{
				[]int{1, 2, 3, 4, 5},
				[]int{16, 17, 18, 19, 6},
				[]int{15, 24, 25, 20, 7},
				[]int{14, 23, 22, 21, 8},
				[]int{13, 12, 11, 10, 9},
			},
		},
		{
			name: "t1",
			args: args{
				n: 3,
			},
			want: [][]int{
				[]int{1, 2, 3},
				[]int{8, 9, 4},
				[]int{7, 6, 5},
			},
		},
		{
			name: "t2",
			args: args{
				n: 1,
			},
			want: [][]int{
				[]int{1},
			},
		},
		{
			name: "t3",
			args: args{
				n: 2,
			},
			want: [][]int{
				[]int{1, 2},
				[]int{4, 3},
			},
		},
		{
			name: "t4",
			args: args{
				n: 4,
			},
			want: [][]int{
				[]int{1, 2, 3, 4},
				[]int{12, 13, 14, 5},
				[]int{11, 16, 15, 6},
				[]int{10, 9, 8, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
