package leetcode

import "testing"

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "t0",
			args: args{
				nums: []int{2, 1},
			},
		},
		{
			name: "t1",
			args: args{
				nums: []int{2, 0, 2, 1, 1, 0},
			},
		},
		{
			name: "t2",
			args: args{
				nums: []int{2, 0, 1},
			},
		},
		{
			name: "t3",
			args: args{
				nums: []int{0},
			},
		},
		{
			name: "t3",
			args: args{
				nums: []int{1},
			},
		},
		{
			name: "t3",
			args: args{
				nums: []int{2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
		})
	}
}
