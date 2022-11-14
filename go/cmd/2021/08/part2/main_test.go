package main

import "testing"

func Test_overlap(t *testing.T) {
	type args struct {
		x []byte
		y []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true",
			args: args{
				x: []byte("ab"),
				y: []byte("abc"),
			},
			want: true,
		},
		{
			name: "should return false",
			args: args{
				x: []byte("acd"),
				y: []byte("ab"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := overlap(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("overlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
