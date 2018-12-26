package main

import "testing"

func TestFindMinTriangleNumberWithMoreThanXFactors(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example",
			args{
				5,
			},
			28,
		},
		{
			"Example",
			args{
				4,
			},
			28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMinTriangleNumberWithMoreThanXFactors(tt.args.x); got != tt.want {
				t.Errorf("FindMinTriangleNumberWithMoreThanXFactors() = %v, want %v", got, tt.want)
			}
		})
	}
}
