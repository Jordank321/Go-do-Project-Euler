package main

import "testing"

func TestSumPrimesToX(t *testing.T) {
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
				10,
			},
			17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumPrimesToX(tt.args.x); got != tt.want {
				t.Errorf("SumPrimesToX() = %v, want %v", got, tt.want)
			}
		})
	}
}
