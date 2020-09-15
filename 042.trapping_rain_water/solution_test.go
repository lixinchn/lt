package leetcode

import "testing"

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			want: 6,
		},
		{
			name: "t2",
			args: args{
				height: []int{0, 5, 0, 5},
			},
			want: 5,
		},
		{
			name: "t3",
			args: args{
				height: []int{0, 5, 5},
			},
			want: 0,
		},
		{
			name: "t4",
			args: args{
				height: []int{1},
			},
			want: 0,
		},
		{
			name: "t5",
			args: args{
				height: []int{7, 1, 0, 6},
			},
			want: 11,
		},
		{
			name: "t6",
			args: args{
				height: []int{7, 7, 1, 7, 1, 0, 6, 9, 1, 6, 4, 9, 5, 8},
			},
			want: 39,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
