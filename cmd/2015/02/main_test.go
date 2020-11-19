package main

import (
	"testing"
)

func Test_wrappingrequired(t *testing.T) {
	type args struct {
		dimensions string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				dimensions: "2x3x4",
			},
			want: 58,
		},
		{
			name: "Example 2",
			args: args{
				dimensions: "1x1x10",
			},
			want: 43,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := wrappingrequired(tt.args.dimensions); got != tt.want || err != nil {
				t.Errorf("wrappingrequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lintrequired(t *testing.T) {
	type args struct {
		dimensions string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "example 1",
			args: args{
				dimensions: "2x3x4",
			},
			want:    34,
			wantErr: false,
		},
		{
			name: "example 2",
			args: args{
				dimensions: "1x1x10",
			},
			want:    14,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lintrequired(tt.args.dimensions)
			if (err != nil) != tt.wantErr {
				t.Errorf("lintrequired() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("lintrequired() = %v, want %v", got, tt.want)
			}
		})
	}
}
