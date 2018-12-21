package main

import (
	"testing"

	"github.com/Jordank321/Go-do-Project-Euler/common"
)

func TestLargestPrimeFactor(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		{
			"Zero",
			args{
				0,
			},
			nil,
		},
		{
			"1 is not prime",
			args{
				1,
			},
			nil,
		},
		{
			"But 2 is a prime",
			args{
				2,
			},
			common.GetIntPointer(2),
		},
		{
			"10",
			args{
				10,
			},
			common.GetIntPointer(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LargestPrimeFactor(tt.args.x)
			if tt.want == nil && got == nil {
				return
			}
			if tt.want == nil && got != nil {
				t.Errorf("Was expecting nil from input %v but found %v", tt.args.x, *got)
				return
			} else if tt.want != nil && got == nil {
				t.Errorf("Was expecting %v from input %v but found nil", *tt.want, tt.args.x)
			} else if *got != *tt.want {
				t.Errorf("LargestPrimeFactor(%v) = %v, want %v", tt.args.x, *got, *tt.want)
			}
		})
	}
}
