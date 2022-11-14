package main

import (
	"testing"
)

func Test_housesvisited(t *testing.T) {
	type args struct {
		directions string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				directions: ">",
			},
			want: 2,
		},
		{
			name: "example 2",
			args: args{
				directions: ">^<v",
			},
			want: 4,
		},
		{
			name: "example 3",
			args: args{
				directions: "^v^v^v^v^v",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := housesvisited(tt.args.directions); got != tt.want {
				t.Errorf("housesvisited() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_housesvisitedwrobotsanta(t *testing.T) {
	type args struct {
		directions string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				directions: "^v",
			},
			want: 3,
		},
		{
			name: "example 2",
			args: args{
				directions: "^>v<",
			},
			want: 3,
		},
		{
			name: "example 3",
			args: args{
				directions: "^v^v^v^v^v",
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := housesvisitedwrobotsanta(tt.args.directions); got != tt.want {
				t.Errorf("housesvisitedwrobotsanta() = %v, want %v", got, tt.want)
			}
		})
	}
}
