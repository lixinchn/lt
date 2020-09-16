package leetcode

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "t1",
			args: args{
				matrix: [][]int{
					[]int{1, 2, 3},
					[]int{4, 5, 6},
					[]int{7, 8, 9},
				},
			},
		},
		{
			name: "t2",
			args: args{
				matrix: [][]int{
					[]int{5, 1, 9, 11},
					[]int{2, 4, 8, 10},
					[]int{13, 3, 6, 7},
					[]int{15, 14, 12, 16},
				},
			},
		},
		{
			name: "t3",
			args: args{
				matrix: [][]int{
					[]int{1, 2},
					[]int{3, 4},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.matrix)
		})
	}
}
