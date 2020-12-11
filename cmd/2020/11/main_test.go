package main

import "testing"

func Test_parseNeighbours(t *testing.T) {
	type args struct {
		hpos int
		vpos int
		plan [][]rune
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "Test 1",
			args: args{
				plan: [][]rune{
					[]rune{'#', '#', '#'},
					[]rune{'#', '#', '#'},
					[]rune{'#', '#', '#'},
				},
				hpos: 1,
				vpos: 0,
			},
			want:  5,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseNeighbours(tt.args.hpos, tt.args.vpos, tt.args.plan)
			if got != tt.want {
				t.Errorf("parseNeighbours() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseNeighbours() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
