package aoc2024

import (
	"testing"
)

func TestIsLevelSafe(t *testing.T) {
	type args struct {
		level []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "safe level 7 6 4 2 1",
			args: args{
				level: []int{7, 6, 4, 2, 1},
			},
			want: true,
		},
		{
			name: "unsafe level 1 2 7 8 9",
			args: args{
				level: []int{1, 2, 7, 8, 9},
			},
			want: false,
		},
		{
			name: "unsafe level 9 7 6 2 1",
			args: args{
				level: []int{9, 7, 6, 2, 1},
			},
			want: false,
		},
		{
			name: "unsafe level 1 3 2 4 5",
			args: args{
				level: []int{1, 3, 2, 4, 5},
			},
			want: false,
		},
		{
			name: "unsafe level 8 6 4 4 1",
			args: args{
				level: []int{8, 6, 4, 4, 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLevelSafe(tt.args.level); got != tt.want {
				t.Errorf("IsLevelSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLevelSafe2(t *testing.T) {
	type args struct {
		level []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "safe level 7 6 4 2 1",
			args: args{
				level: []int{7, 6, 4, 2, 1},
			},
			want: true,
		},
		{
			name: "unsafe level 1 2 7 8 9",
			args: args{
				level: []int{1, 2, 7, 8, 9},
			},
			want: false,
		},
		{
			name: "unsafe level 9 7 6 2 1",
			args: args{
				level: []int{9, 7, 6, 2, 1},
			},
			want: false,
		},
		{
			name: "unsafe level 1 3 2 4 5",
			args: args{
				level: []int{1, 3, 2, 4, 5},
			},
			want: true,
		},
		{
			name: "unsafe level 8 6 4 4 1",
			args: args{
				level: []int{8, 6, 4, 4, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLevelSafe2(tt.args.level); got != tt.want {
				t.Errorf("IsLevelSafe2() = %v, want %v", got, tt.want)
			}
		})
	}
}
