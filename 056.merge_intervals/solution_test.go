package leetcode

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "t1",
			args: args{
				intervals: [][]int{
					[]int{2, 6},
					[]int{1, 3},
					[]int{8, 10},
					[]int{15, 18},
				},
			},
			want: [][]int{
				[]int{1, 6},
				[]int{8, 10},
				[]int{15, 18},
			},
		},
		{
			name: "t2",
			args: args{
				intervals: [][]int{
					[]int{1, 4},
					[]int{4, 5},
				},
			},
			want: [][]int{
				[]int{1, 5},
			},
		},
		{
			name: "t3",
			args: args{
				intervals: [][]int{
					[]int{1, 10},
					[]int{4, 5},
					[]int{8, 9},
				},
			},
			want: [][]int{
				[]int{1, 10},
			},
		},
		{
			name: "t4",
			args: args{
				intervals: [][]int{
					[]int{1, 10},
				},
			},
			want: [][]int{
				[]int{1, 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
