package aoc2023

import (
	"reflect"
	"testing"
)

func TestReadLine(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want Game
	}{
		{
			name: "success",
			args: args{
				input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			},
			want: Game{
				id: 1,
				sets: []Set{
					{
						Blue: 3,
						Red:  4,
					},
					{
						Red:   1,
						Green: 2,
						Blue:  6,
					},
					{
						Green: 2,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadLine(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_IsValid(t *testing.T) {
	type fields struct {
		id   int
		sets []Set
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid",
			fields: fields{
				sets: []Set{{
					Blue: 14,
				}}},
			want: true,
		},
		{
			name: "valid",
			fields: fields{
				sets: []Set{{
					Red:   1,
					Green: 1,
					Blue:  1,
				}}},
			want: true,
		},
		{
			name: "invalid",
			fields: fields{
				sets: []Set{{
					Red:   15,
					Green: 16,
					Blue:  16,
				}}},
			want: false,
		},
		{
			name: "invalid",
			fields: fields{
				sets: []Set{{
					Red:   122,
					Green: 1,
					Blue:  1,
				}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Game{
				id:   tt.fields.id,
				sets: tt.fields.sets,
			}
			if got := g.IsValid(); got != tt.want {
				t.Errorf("Game.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveCubeConundrum(t *testing.T) {
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
				input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveCubeConundrum(tt.args.input); got != tt.want {
				t.Errorf("SolveCubeConundrum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveCubeConundrum2(t *testing.T) {
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
				input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			want: 2286,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveCubeConundrum2(tt.args.input); got != tt.want {
				t.Errorf("SolveCubeConundrum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
