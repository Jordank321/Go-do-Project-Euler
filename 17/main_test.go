package main

import (
	"testing"
)

func TestCountCharacters(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"This is a test!",
			args{
				"This is a test-test",
			},
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountCharacters(tt.args.in); got != tt.want {
				t.Errorf("CountCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNumberAsWords(t *testing.T) {
	type args struct {
		in int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"One hundred and twelve",
			args{
				112,
			},
			"onehundredandtwelve",
		},
		{
			"three hundred and forty-two",
			args{
				342,
			},
			"threehundredandfortytwo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNumberAsWords(tt.args.in); got != tt.want {
				t.Errorf("GetNumberAsWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
