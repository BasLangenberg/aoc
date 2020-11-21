package main

import "testing"

func Test_isNice(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				line: "ugknbfddgicrmopn",
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				line: "aaa",
			},
			want: true,
		},
		{
			name: "example 3",
			args: args{
				line: "jchzalrnumimnmhp",
			},
			want: false,
		},
		{
			name: "example 4",
			args: args{
				line: "haegwjzuvuyypxyu",
			},
			want: false,
		},
		{
			name: "example 5",
			args: args{
				line: "dvszwmarrgswjxmb",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNice(tt.args.line); got != tt.want {
				t.Errorf("isNice() = %v, want %v", got, tt.want)
			}
		})
	}
}
