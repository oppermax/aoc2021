package main

import "testing"

func Test_alignAtComplex(t *testing.T) {
	type args struct {
		level int
		in    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "one",
			args: args{
				level: 0,
				in:    []int{5},
			},
		want: 15},
		{name: "one",
			args: args{
				level: 5,
				in:    []int{2},
			},
			want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alignAtComplex(tt.args.level, tt.args.in); got != tt.want {
				t.Errorf("alignAtComplex() = %v, want %v", got, tt.want)
			}
		})
	}
}
