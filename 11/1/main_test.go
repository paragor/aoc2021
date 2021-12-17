package main

import (
	"reflect"
	"testing"
)

func Test_getPontsFor(t *testing.T) {
	type args struct {
		field [][]*Cell
		y     int
		x     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "simple", args: args{
			field: [][]*Cell{
				{nil, nil, nil},
				{nil, nil, nil},
				{nil, nil, nil},
			},
			y:     1,
			x:     1,
		}, want: [][]int{
			{2, 1},
			{2, 2},
			{2, 0},
			{0, 1},
			{0, 2},
			{0, 0},
			{1, 2},
			{1, 0},
		}},
		{name: "simple", args: args{
			field: [][]*Cell{
				{nil, nil, nil},
				{nil, nil, nil},
				{nil, nil, nil},
			},
			y:     0,
			x:     0,
		}, want: [][]int{
			{1, 0},
			{1, 1},
			{0, 1},
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPontsFor(tt.args.field, tt.args.y, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPontsFor() = %v, want %v", got, tt.want)
			}
		})
	}
}
