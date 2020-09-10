package leetcode

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "t1",
			args: args{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			want: [][]int{
				[]int{-1, 0, 0, 1},
				[]int{-2, -1, 1, 2},
				[]int{-2, 0, 0, 2},
			},
		},
		{
			name: "t2",
			args: args{
				nums:   []int{-1, 0, 1, 2, -1, -4},
				target: -1,
			},
			want: [][]int{
				[]int{-4, 0, 1, 2},
				[]int{-1, -1, 0, 1},
			},
		},
		{
			name: "t3",
			args: args{
				nums:   []int{-1, 0, -5, -2, -2, -4, 0, 1, -2},
				target: -9,
			},
			want: [][]int{
				[]int{-5, -4, -1, 1},
				[]int{-5, -4, 0, 0},
				[]int{-5, -2, -2, 0},
				[]int{-4, -2, -2, -1},
			},
		},
		{
			name: "t4",
			args: args{
				nums:   []int{0, 0, 0, 0},
				target: 0,
			},
			want: [][]int{
				[]int{0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
