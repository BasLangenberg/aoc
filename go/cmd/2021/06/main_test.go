package main

import (
	"testing"
)

func Test_simulateRepro(t *testing.T) {
	type args struct {
		fish []int
		days int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Part one: AOC Example",
			args: args{
				fish: []int{3, 4, 3, 1, 2},
				days: 80,
			},
			want: 5934,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simulateRepro(tt.args.fish, tt.args.days); got != tt.want {
				t.Errorf("simulateRepro() = %v, want %v", got, tt.want)
			}
		})
	}
}
