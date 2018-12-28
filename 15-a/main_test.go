package main

import "testing"

func TestFindRoutesInSquareGridOfX(t *testing.T) {
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
				2,
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindRoutesInSquareGridOfX(tt.args.x); got != tt.want {
				t.Errorf("FindRoutesInSquareGridOfX() = %v, want %v", got, tt.want)
			}
		})
	}
}
