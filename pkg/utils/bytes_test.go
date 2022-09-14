package utils

import (
	"reflect"
	"testing"
)

func Test_anyIntToBytes(t *testing.T) {
	type args struct {
		x    uint64
		size int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "test zero", args: args{x: 0, size: 64}, want: []byte{0, 0, 0, 0, 0, 0, 0, 0}},
		{name: "test one as 64", args: args{x: 1, size: 64}, want: []byte{0, 0, 0, 0, 0, 0, 0, 1}},
		{name: "test one as 63", args: args{x: 1, size: 64}, want: []byte{0, 0, 0, 0, 0, 0, 0, 1}},
		{name: "test 1010000000000000 as 16", args: args{x: 0b1010000000000000}, want: []byte{0b10100000, 0}},
		{name: "test 1010000000000000 as 17", args: args{x: 0b1010000000000000}, want: []byte{0, 0b10100000, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnyNumberToBytes(tt.args.x, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnyNumberToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolArrayToBytes(t *testing.T) {
	type args struct {
		bs []bool
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "test - 1", args: args{[]bool{true, true, true, true}}, want: []byte{0b00001111}},
		{name: "test - 2", args: args{[]bool{true, true, true, true, true, true, true, true}}, want: []byte{0b11111111}},
		{name: "test - 3", args: args{[]bool{true, true, true, true, true, true, true, true, true}}, want: []byte{0b00000001, 0b11111111}},
		{name: "test - 4", args: args{[]bool{true, false, true, false, true, false, true, false, true}}, want: []byte{0b00000001, 0b01010101}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolArrayToBytes(tt.args.bs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoolArrayToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
