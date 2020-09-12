package leetcode

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "t1",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{
				[]int{7},
				[]int{2, 2, 3},
			},
		},
		{
			name: "t2",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			want: [][]int{
				[]int{2, 2, 2, 2},
				[]int{2, 3, 3},
				[]int{3, 5},
			},
		},
		{
			name: "t3",
			args: args{
				candidates: []int{7, 3, 2},
				target:     18,
			},
			want: [][]int{
				[]int{2, 2, 2, 2, 2, 2, 2, 2, 2},
				[]int{2, 2, 2, 2, 2, 2, 3, 3},
				[]int{2, 2, 2, 2, 3, 7},
				[]int{2, 2, 2, 3, 3, 3, 3},
				[]int{2, 2, 7, 7},
				[]int{2, 3, 3, 3, 7},
				[]int{3, 3, 3, 3, 3, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
