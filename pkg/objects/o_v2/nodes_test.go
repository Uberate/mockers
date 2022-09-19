package o_v2

import (
	"reflect"
	"testing"
)

func TestNodeData_To(t *testing.T) {
	type args struct {
		obj struct {
			Test string `mapstructure:"test"`
		}
	}
	tests := []struct {
		name string
		n    NodeData
		args args
		wantErr bool
	}{
		{name: "test node data to object", n: NodeData{"test": "test"}, args: args{obj: struct {
			Test string `mapstructure:"test"`
		}{}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.To(&tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("To() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				println(tt.args.obj.Test)
			}
		})
	}
}

func TestNewNodeData(t *testing.T) {
	type args struct {
		obj struct {
			Test string `mapstructure:"test"`
		}
	}
	tests := []struct {
		name    string
		args    args
		want    NodeData
		wantErr bool
	}{
		{name: "test object to node data", args: args{obj: struct {
			Test string `mapstructure:"test"`
		}{Test: "test"}}, want: NodeData{"test": "test"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewNodeData(tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewNodeData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNodeData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
