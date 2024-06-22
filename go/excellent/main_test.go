package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Even number",
			args: args{number: 8},
			want: "Even",
		},
		{
			name: "Odd number",
			args: args{number: 11},
			want: "Odd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvenOrOdd(tt.args.number); got != tt.want {
				t.Errorf("EvenOrOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
