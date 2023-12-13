package aoc2023

import "testing"

func Test_calibrateValue(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base case",
			args: args{
				input: "pqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet",
			},
			want: 142,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calibrateValue(tt.args.input); got != tt.want {
				t.Errorf("calibrateValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
