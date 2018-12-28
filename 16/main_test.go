package main

import "testing"

func TestGetSumOfDigitsOfXToThePowerOfY(t *testing.T) {
	type args struct {
		x int64
		y int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Given example",
			args{
				2,
				15,
			},
			26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSumOfDigitsOfXToThePowerOfY(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("GetSumOfDigitsOfXToThePowerOfY() = %v, want %v", got, tt.want)
			}
		})
	}
}
