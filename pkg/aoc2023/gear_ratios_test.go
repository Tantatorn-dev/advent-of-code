package aoc2023

import "testing"

func Test_SumPartNumbers(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				input: "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..",
			},
			want: 4361,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumPartNumbers(tt.args.input); got != tt.want {
				t.Errorf("sumPartNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
