package aoc2023

import (
	"testing"
)

func TestSumExtrapolatedValues(t *testing.T) {
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
				input: "0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45",
			},
			want: 114,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumExtrapolatedValues(tt.args.input); got != tt.want {
				t.Errorf("SumExtrapolatedValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindExtrapolatedValue(t *testing.T) {
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
				input: "10 13 16 21 30 45",
			},
			want: 68,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findExtrapolatedValue(tt.args.input); got != tt.want {
				t.Errorf("FindExtrapolatedValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
