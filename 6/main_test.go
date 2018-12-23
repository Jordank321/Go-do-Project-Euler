package main

import (
	"testing"
)

func TestFindSumOfSquaresToX(t *testing.T) {
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
			385,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSumOfSquaresToX(tt.args.x); got != tt.want {
				t.Errorf("FindSumOfSquaresToX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindSquareOfSumsToX(t *testing.T) {
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
			3025,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSquareOfSumsToX(tt.args.x); got != tt.want {
				t.Errorf("FindSquareOfSumsToX() = %v, want %v", got, tt.want)
			}
		})
	}
}
