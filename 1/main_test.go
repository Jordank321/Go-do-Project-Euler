package main

import "testing"

func TestGetSumOfMultiplesOf3And5LessThanX(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Given example",
			args{
				10,
			},
			23,
		},
		{
			"Zero",
			args{
				0,
			},
			0,
		},
		{
			"Multiples from both",
			args{
				31,
			},
			225,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSumOfMultiplesOf3And5LessThanX(tt.args.x); got != tt.want {
				t.Errorf("GetSumOfMultiplesOf3And5LessThanX(%v) = %v, want %v", tt.args.x, got, tt.want)
			}
		})
	}
}
