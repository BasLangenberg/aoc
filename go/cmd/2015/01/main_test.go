package main

import (
	"testing"
)

func Test_getfloor(t *testing.T) {
	type args struct {
		in0 []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				in0: []byte("(())"),
			},
			want: 0,
		},
		{
			name: "Example 2",
			args: args{
				in0: []byte("()()"),
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				in0: []byte("))((((("),
			},
			want: 3,
		},
		{
			name: "Example 4",
			args: args{
				in0: []byte("))("),
			},
			want: -1,
		},
		{
			name: "Example 5",
			args: args{
				in0: []byte(")())())"),
			},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getfloor(tt.args.in0); got != tt.want {
				t.Errorf("getfloor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basementhit(t *testing.T) {
	type args struct {
		instructions []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				instructions: []byte(")"),
			},
			want: 1,
		},
		{
			name: "Example 1",
			args: args{
				instructions: []byte("()())"),
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basementhit(tt.args.instructions); got != tt.want {
				t.Errorf("basementhit() = %v, want %v", got, tt.want)
			}
		})
	}
}
