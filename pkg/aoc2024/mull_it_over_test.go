package aoc2024

import (
	"testing"
)

func TestMultiplyByInstruction(t *testing.T) {
	type args struct {
		instruction string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				instruction: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			},
			want: 161,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplyByInstruction(tt.args.instruction); got != tt.want {
				t.Errorf("MultiplyByInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplyByInstruction2(t *testing.T) {
	type args struct {
		instruction string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				instruction: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			},
			want: 48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplyByInstruction2(tt.args.instruction); got != tt.want {
				t.Errorf("MultiplyByInstruction2() = %v, want %v", got, tt.want)
			}
		})
	}
}
