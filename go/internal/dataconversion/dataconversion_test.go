// Package dataconversion is written for use with AOC problems only
// Therefore, I do not think it is an antipattern to terminate inside this package
// Huge exception to normal coding standards!
package dataconversion

import (
	"reflect"
	"testing"
)

func TestStringToIntArray(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				input: "1,2,4",
			},
			want: []int{1, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToIntArray(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToIntArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
