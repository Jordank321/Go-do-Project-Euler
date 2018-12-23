package main

import "testing"

func TestFindPrimeX(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example given",
			args{
				6,
			},
			13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPrimeX(tt.args.x); got != tt.want {
				t.Errorf("FindPrimeX() = %v, want %v", got, tt.want)
			}
		})
	}
}
