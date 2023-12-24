package aoc2023

import (
	"reflect"
	"testing"
)

func Test_getMap(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "success",
			args: args{
				input: "50 98 2\n52 10 2",
			},
			want: map[int]int{
				98: 50,
				99: 51,
				10: 52,
				11: 53,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMap(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLowestLocation(t *testing.T) {
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
				input: `seeds: 79 14 55 13

				seed-to-soil map:
				50 98 2
				52 50 48
				
				soil-to-fertilizer map:
				0 15 37
				37 52 2
				39 0 15
				
				fertilizer-to-water map:
				49 53 8
				0 11 42
				42 0 7
				57 7 4
				
				water-to-light map:
				88 18 7
				18 25 70
				
				light-to-temperature map:
				45 77 23
				81 45 19
				68 64 13
				
				temperature-to-humidity map:
				0 69 1
				1 0 69
				
				humidity-to-location map:
				60 56 37
				56 93 4`,
			},
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLowestLocation(tt.args.input); got != tt.want {
				t.Errorf("GetLowestLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
