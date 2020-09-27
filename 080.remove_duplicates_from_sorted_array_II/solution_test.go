package leetcode

import "testing"

func Test_removeDuplicates(t *testing.T) {
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
				nums: []int{1, 2, 2, 2, 2, 2, 2},
			},
			want: 3,
		},
		{
			name: "t2",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 3,
		},
		{
			name: "t3",
			args: args{
				nums: []int{1, 2, 2, 2, 3},
			},
			want: 4,
		},
		{
			name: "t4",
			args: args{
				nums: []int{1, 2, 2, 2, 2, 3},
			},
			want: 4,
		},
		{
			name: "t5",
			args: args{
				nums: []int{1, 2, 2, 2, 2, 2, 3, 3},
			},
			want: 5,
		},
		{
			name: "t6",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			},
			want: 7,
		},
		{
			name: "t7",
			args: args{
				nums: []int{1, 1, 1, 1, 2, 2, 3},
			},
			want: 5,
		},
		{
			name: "t8",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
			},
			want: 5,
		},
		{
			name: "t9",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 2, 3, 3},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
