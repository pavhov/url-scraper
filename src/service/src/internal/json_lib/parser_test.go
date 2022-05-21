package json_lib

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	type IType map[string]any
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want IType
	}{
		{
			name: "",
			args: args{
				data: `{"val_1":"test1"}`,
			},
			want: map[string]any{
				"val_1": "test1",
			},
		},
		{
			name: "",
			args: args{
				data: "",
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				data: `{"val_1":"test1}`,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode[IType](tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	type IType map[string]any
	type args struct {
		data IType
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				data: map[string]any{
					"val_1": "test1",
				},
			},
			want: `{"val_1":"test1"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode[IType](tt.args.data); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
