package leetcode

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},
		{
			name: "t2",
			args: args{
				nums:   []int{-1, 2, 2, 2, 2, 2, 2, 2, 1, -4},
				target: 1,
			},
			want: 0,
		},
		{
			name: "t3",
			args: args{
				nums:   []int{-1, 2, 2, 2, 2, 2, 2, 2, 1, -4},
				target: 0,
			},
			want: 0,
		},
		{
			name: "t4",
			args: args{
				nums:   []int{1, 2, 4, 8, 16, 32, 64, 128},
				target: 82,
			},
			want: 82,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
