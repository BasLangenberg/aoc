package main

import "testing"

func TestLeastFuel(t *testing.T) {
	tstInput := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	if got := leastfuel(tstInput); got != 37 {
		t.Errorf("simulateRepro() = %v, want %v", got, 37)
	}

}

func TestLeastFuel2(t *testing.T) {
	tstInput := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	if got := leastfuel2(tstInput); got != 168 {
		t.Errorf("simulateRepro() = %v, want %v", got, 168)
	}

}
