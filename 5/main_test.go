package main

import "testing"

func TestFindSmallestDivisiableByAllUpToX(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Given Example",
			args{
				10,
			},
			2520,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSmallestDivisiableByAllUpToX(tt.args.x); got != tt.want {
				t.Errorf("FindSmallestDivisiableByAllUpToX() = %v, want %v", got, tt.want)
			}
		})
	}
}
