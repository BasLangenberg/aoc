package main

import "testing"

func Test_csasasalcHash(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				key: "abcdef",
			},
			want: 609043,
		},
		{
			name: "example 1",
			args: args{
				key: "pqrstuv",
			},
			want: 1048970,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcHash(tt.args.key); got != tt.want {
				t.Errorf("calcHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
