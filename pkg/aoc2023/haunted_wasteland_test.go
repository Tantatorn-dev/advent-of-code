package aoc2023

import (
	"reflect"
	"testing"
)

func Test_constructWastelandMap(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{
			name: "construct map from string success",
			args: args{
				"aaa = (bbb, ccc)\nbbb = (ddd, eee)\nccc = (fff, ggg)\n",
			},
			want: map[string][]string{
				"aaa": {"bbb", "ccc"},
				"bbb": {"ddd", "eee"},
				"ccc": {"fff", "ggg"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructWastelandMap(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructWastelandMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTraverseWasteland(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2 steps to ZZZ",
			args: args{
				input: "RL\n\nAAA = (BBB, CCC)\nBBB = (DDD, EEE)\nCCC = (ZZZ, GGG)\nDDD = (DDD, DDD)\nEEE = (EEE, EEE)\nGGG = (GGG, GGG)\nZZZ = (ZZZ, ZZZ)",
			},
			want: 2,
		},
		{
			name: "6 steps to ZZZ",
			args: args{
				input: "LLR\n\nAAA = (BBB, BBB)\nBBB = (AAA, ZZZ)\nZZZ = (ZZZ, ZZZ)",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TraverseWasteland(tt.args.input); got != tt.want {
				t.Errorf("TraverseWasteland() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTraverseWasteland2(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Base case",
			args: args{
				input: "LR\n\n11A = (11B, XXX)\n11B = (XXX, 11Z)\n11Z = (11B, XXX)\n22A = (22B, XXX)\n22B = (22C, 22C)\n22C = (22Z, 22Z)\n22Z = (22B, 22B)\nXXX = (XXX, XXX)",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TraverseWasteland2(tt.args.input); got != tt.want {
				t.Errorf("TraverseWasteland2() = %v, want %v", got, tt.want)
			}
		})
	}
}
