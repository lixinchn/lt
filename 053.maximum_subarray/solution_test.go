package leetcode

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "t2",
			args: args{
				nums: []int{-2, -1},
			},
			want: -1,
		},
		{
			name: "t3",
			args: args{
				nums: []int{-2147483647},
			},
			want: -2147483647,
		},
		{
			name: "t4",
			args: args{
				nums: []int{-1, -2, -3, -4},
			},
			want: -1,
		},
		{
			name: "t5",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 10,
		},
		{
			name: "t6",
			args: args{
				nums: []int{0, 0, 2, -1, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
