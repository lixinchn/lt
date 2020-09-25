package leetcode

import "testing"

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "t1",
			args: args{
				board: [][]byte{
					[]byte{'C', 'A', 'A'},
					[]byte{'A', 'A', 'A'},
					[]byte{'B', 'C', 'D'},
				},
				word: "AAB",
			},
			want: true,
		},
		{
			name: "t1",
			args: args{
				board: [][]byte{
					[]byte{'A', 'A'},
				},
				word: "AA",
			},
			want: true,
		},
		{
			name: "t1",
			args: args{
				board: [][]byte{
					[]byte{'A', 'B', 'C', 'E'},
					[]byte{'S', 'F', 'C', 'S'},
					[]byte{'A', 'D', 'E', 'E'},
				},
				word: "ABCB",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
