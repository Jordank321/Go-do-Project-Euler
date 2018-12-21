package main

import "testing"

func TestSumFibonacciBelowX(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Zero",
			args{
				0,
			},
			0,
		},
		{
			"First",
			args{
				2,
			},
			0,
		},
		{
			"Second",
			args{
				3,
			},
			2,
		},
		{
			"Up to 12",
			args{
				13,
			},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumEvenFibonacciBelowX(tt.args.x); got != tt.want {
				t.Errorf("SumEvenFibonacciBelowX(%v) = %v, want %v", tt.args.x, got, tt.want)
			}
		})
	}
}
