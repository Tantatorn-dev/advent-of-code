package aoc2024

import "testing"

func TestFindTotalDistances(t *testing.T) {
	type args struct {
		l1 []int64
		l2 []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "success",
			args: args{
				l1: []int64{1, 2, 3},
				l2: []int64{7, 8, 9},
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTotalDistances(tt.args.l1, tt.args.l2); got != tt.want {
				t.Errorf("FindTotalDistances() = %v, want %v", got, tt.want)
			}
		})
	}
}
