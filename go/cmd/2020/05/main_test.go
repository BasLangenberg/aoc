package main

import "testing"

func Test_checkSeat(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name   string
		args   args
		row    int
		column int
		seatid int
	}{
		{
			name: "FBFBBFFRLR",
			args: args{
				input: "FBFBBFFRLR",
			},
			row:    44,
			column: 5,
			seatid: 357,
		},
		{
			name: "BFFFBBFRRR",
			args: args{
				input: "BFFFBBFRRR",
			},
			row:    70,
			column: 7,
			seatid: 567,
		},
		{
			name: "FFFBBBFRRR",
			args: args{
				input: "FFFBBBFRRR",
			},
			row:    14,
			column: 7,
			seatid: 119,
		},
		{
			name: "BBFFBBFRLL",
			args: args{
				input: "BBFFBBFRLL",
			},
			row:    102,
			column: 4,
			seatid: 820,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := checkSeat(tt.args.input)
			if got != tt.row {
				t.Errorf("checkSeat() got = %v, want %v", got, tt.row)
			}
			if got1 != tt.column {
				t.Errorf("checkSeat() got1 = %v, want %v", got1, tt.column)
			}
			if got2 != tt.seatid {
				t.Errorf("checkSeat() got2 = %v, want %v", got2, tt.seatid)
			}
		})
	}
}
