package main

import "testing"

func TestFindLargestAdjacentNumbers(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Small grid",
			args{
				`9 2 3 7 7
				4 9 6 7 7
				7 8 9 7 7
				7 7 7 9 7`,
			},
			6561,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLargestAdjacentNumbers(tt.args.input); got != tt.want {
				t.Errorf("FindLargestAdjacentNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
