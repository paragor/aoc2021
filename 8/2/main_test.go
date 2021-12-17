package main

import (
	"reflect"
	"testing"
)

func TestWire(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "origin",
			args: args{strs: []string{"cf", "acf", "bcdf", "abdfg", "acdeg", "acdfg", "abcefg", "abdefg", "abcdfg", "abcdefg"}},
			want: map[string]string{
				"a": "a",
				"b": "b",
				"c": "c",
				"d": "d",
				"e": "e",
				"f": "f",
				"g": "g",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wire(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wire() = %v, want %v", got, tt.want)
			}
		})
	}
}
