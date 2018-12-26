package main

import "testing"

func TestGetFirst10DigitsFromSumOf(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Simple example",
			args{
				`11111111111
11111111111`,
			},
			"2222222222",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFirst10DigitsFromSumOf(tt.args.input); got != tt.want {
				t.Errorf("GetFirst10DigitsFromSumOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
