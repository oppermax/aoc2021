package main

import (
	"reflect"
	"testing"
)

func Test_getDiagonalPositions(t *testing.T) {
	type args struct {
		a position
		b position
	}
	tests := []struct {
		name string
		args args
		want []position
	}{
		{name: "test-1",
			args: args{
				a: position{
					x: 8,
					y: 0,
				},
				b: position{
					x: 0,
					y: 8,
				},
			}},
		{name: "test-2",
			args: args{
				a: position{
					x: 6,
					y: 4,
				},
				b: position{
					x: 2,
					y: 0,
				},
			}},
		{name: "test-3",
			args: args{
				a: position{
					x: 0,
					y: 9,
				},
				b: position{
					x: 5,
					y: 9,
				},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDiagonalPositions(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDiagonalPositions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createVent(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want vent
	}{
		{name: "test-1",
			args: args{in: "9,4 -> 3,4"},
		want: vent{
			start: position{
				x: 3,
				y: 4,
			},
			end:   position{
				x: 9,
				y: 4,
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createVent(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createVent() = %v, want %v", got, tt.want)
			}
		})
	}
}