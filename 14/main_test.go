package main

import "testing"

func TestFindLongestChainUnderX(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example from wikipidia",
			args{
				10,
			},
			9,
		},
		{
			"Example from wikipidia",
			args{
				100,
			},
			97,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLongestChainUnderX(tt.args.x); got != tt.want {
				t.Errorf("FindLongestChainUnderX() = %v, want %v", got, tt.want)
			}
		})
	}
}
