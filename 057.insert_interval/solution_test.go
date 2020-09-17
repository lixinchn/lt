package leetcode

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "t-6",
			args: args{
				intervals: [][]int{
					[]int{1, 4},
					[]int{9, 12},
					[]int{19, 22},
				},
				newInterval: []int{7, 13},
			},
			want: [][]int{
				[]int{1, 4},
				[]int{7, 13},
				[]int{19, 22},
			},
		},
		{
			name: "t-5",
			args: args{
				intervals: [][]int{
					[]int{0, 10},
					[]int{14, 14},
					[]int{15, 20},
				},
				newInterval: []int{11, 11},
			},
			want: [][]int{
				[]int{0, 10},
				[]int{11, 11},
				[]int{14, 14},
				[]int{15, 20},
			},
		},
		{
			name: "t-4",
			args: args{
				intervals: [][]int{
					[]int{3, 5},
					[]int{12, 15},
				},
				newInterval: []int{6, 6},
			},
			want: [][]int{
				[]int{3, 5},
				[]int{6, 6},
				[]int{12, 15},
			},
		},
		{
			name: "t-3",
			args: args{
				intervals: [][]int{
					[]int{0, 5},
					[]int{9, 12},
				},
				newInterval: []int{7, 16},
			},
			want: [][]int{
				[]int{0, 5},
				[]int{7, 16},
			},
		},
		{
			name: "t-2",
			args: args{
				intervals: [][]int{
					[]int{2, 6},
					[]int{7, 9},
				},
				newInterval: []int{15, 18},
			},
			want: [][]int{
				[]int{2, 6},
				[]int{7, 9},
				[]int{15, 18},
			},
		},
		{
			name: "t-1",
			args: args{
				intervals: [][]int{
					[]int{0, 2},
					[]int{3, 3},
					[]int{6, 11},
				},
				newInterval: []int{9, 15},
			},
			want: [][]int{
				[]int{0, 2},
				[]int{3, 3},
				[]int{6, 15},
			},
		},
		{
			name: "t0",
			args: args{
				intervals: [][]int{
					[]int{1, 5},
				},
				newInterval: []int{6, 8},
			},
			want: [][]int{
				[]int{1, 5},
				[]int{6, 8},
			},
		},
		{
			name: "t0.1",
			args: args{
				intervals: [][]int{
					[]int{1, 5},
				},
				newInterval: []int{5, 8},
			},
			want: [][]int{
				[]int{1, 8},
			},
		},
		{
			name: "t0.2",
			args: args{
				intervals: [][]int{
					[]int{5, 8},
				},
				newInterval: []int{1, 2},
			},
			want: [][]int{
				[]int{1, 2},
				[]int{5, 8},
			},
		},
		{
			name: "t0.3",
			args: args{
				intervals: [][]int{
					[]int{5, 8},
				},
				newInterval: []int{1, 5},
			},
			want: [][]int{
				[]int{1, 8},
			},
		},
		{
			name: "t0.4",
			args: args{
				intervals: [][]int{
					[]int{1, 5},
				},
				newInterval: []int{0, 3},
			},
			want: [][]int{
				[]int{0, 5},
			},
		},
		{
			name: "t0.5",
			args: args{
				intervals: [][]int{
					[]int{1, 5},
				},
				newInterval: []int{0, 6},
			},
			want: [][]int{
				[]int{0, 6},
			},
		},
		{
			name: "t0.6",
			args: args{
				intervals: [][]int{
					[]int{1, 5},
					[]int{6, 8},
				},
				newInterval: []int{5, 6},
			},
			want: [][]int{
				[]int{1, 8},
			},
		},
		{
			name: "t0.7",
			args: args{
				intervals: [][]int{
					[]int{0, 5},
					[]int{8, 9},
				},
				newInterval: []int{3, 4},
			},
			want: [][]int{
				[]int{0, 5},
				[]int{8, 9},
			},
		},
		{
			name: "t1",
			args: args{
				intervals: [][]int{
					[]int{1, 3},
					[]int{6, 9},
				},
				newInterval: []int{2, 5},
			},
			want: [][]int{
				[]int{1, 5},
				[]int{6, 9},
			},
		},
		{
			name: "t2",
			args: args{
				intervals: [][]int{
					[]int{1, 2},
					[]int{3, 5},
					[]int{6, 7},
					[]int{8, 10},
					[]int{12, 16},
				},
				newInterval: []int{4, 8},
			},
			want: [][]int{
				[]int{1, 2},
				[]int{3, 10},
				[]int{12, 16},
			},
		},
		{
			name: "t3",
			args: args{
				intervals:   [][]int{},
				newInterval: []int{5, 7},
			},
			want: [][]int{
				[]int{5, 7},
			},
		},
		{
			name: "t4",
			args: args{
				intervals: [][]int{
					[]int{1, 5},
				},
				newInterval: []int{2, 3},
			},
			want: [][]int{
				[]int{1, 5},
			},
		},
		{
			name: "t5",
			args: args{
				intervals: [][]int{
					[]int{1, 5},
				},
				newInterval: []int{2, 7},
			},
			want: [][]int{
				[]int{1, 7},
			},
		},
		{
			name: "t6",
			args: args{
				intervals: [][]int{
					[]int{3, 5},
					[]int{8, 10},
				},
				newInterval: []int{1, 17},
			},
			want: [][]int{
				[]int{1, 17},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
