package aoc2023

import "testing"

func Test_determineCardHand(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want CardHand
	}{
		{
			name: "five of a kind",
			args: args{
				input: "11111",
			},
			want: FiveOfAKind,
		},
		{
			name: "four of a kind",
			args: args{
				input: "aaaab",
			},
			want: FourOfAKind,
		},
		{
			name: "full house",
			args: args{
				input: "aaabb",
			},
			want: FullHouse,
		},
		{
			name: "two pair",
			args: args{
				input: "aaddx",
			},
			want: TwoPair,
		},
		{
			name: "one pair",
			args: args{
				input: "aaxyz",
			},
			want: OnePair,
		},
		{
			name: "high card",
			args: args{
				input: "abcde",
			},
			want: HighCard,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determineCardHand(tt.args.input); got != tt.want {
				t.Errorf("isFiveOfAKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTotalWinning(t *testing.T) {
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
				input: `32T3K 765
						T55J5 684
						KK677 28
						KTJJT 220
						QQQJA 483`,
			},
			want: 6440,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTotalWinning(tt.args.input); got != tt.want {
				t.Errorf("GetTotalWinning() = %v, want %v", got, tt.want)
			}
		})
	}
}
