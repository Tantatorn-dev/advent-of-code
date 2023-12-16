package aoc2023

import (
	"testing"
)

func Test_CalibrateValue(t *testing.T) {
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
				input: "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet",
			},
			want: 142,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalibrateValue(tt.args.input); got != tt.want {
				t.Errorf("calibrateValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalibrateValue2(t *testing.T) {
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
				input: "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen",
			},
			want: 281,
		},
		{
			name: "success",
			args: args{
				input: "two\neightwothree",
			},
			want: 105,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalibrateValue2(tt.args.input); got != tt.want {
				t.Errorf("CalibrateValue2() = %v, want %v", got, tt.want)
			}
		})
	}
}
