package main

import (
	"bufio"
	"os"
	"testing"
)

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

func Test_isNiceTwoPointZero(t *testing.T) {
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
				line: "qjhvhtzxzqqjkmpb",
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				line: "xxyxx",
			},
			want: true,
		},
		{
			name: "example 3",
			args: args{
				line: "uurcxstgmygtbstg",
			},
			want: false,
		},
		{
			name: "example 4",
			args: args{
				line: "ieodomkazucvgmuy",
			},
			want: false,
		},
		{
			name: "example 5",
			args: args{
				line: "rxexcbwhiywwwwnu",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNiceTwoPointZero(tt.args.line); got != tt.want {
				t.Errorf("isNiceTwoPointZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Debug
// cat input |  grep "\(..\).*\1" | grep "\(.\).\1" | wc -l Finds all the incorrect matches
func Test_DebugNotFoundString(t *testing.T) {
	input, err := os.Open("naughty2point0strings")
	defer input.Close()

	if err != nil {
		t.Fatalf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	inputs := make(map[string]bool)

	for scanner.Scan() {
		line := scanner.Text()
		nice := isNiceTwoPointZero(line)

		inputs[line] = nice
	}

	for line, result := range inputs {
		if !result {
			t.Errorf("%v should be nice but is not nice\n", line)
		}
	}
}
