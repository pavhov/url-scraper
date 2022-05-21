package api

import (
	"reflect"
	"service/src/internal/http_lib"
	"testing"
)

func TestApi_Init(t *testing.T) {
	type fields struct {
		Http http_lib.Interface
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Api_Init1",
			fields: fields{
				Http: http_lib.NewServer(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Api{
				Http: tt.fields.Http,
			}
			if err := a.Init(); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApi_defineRouters(t *testing.T) {
	type fields struct {
		Http http_lib.Interface
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Api_defineRouters1",
			fields: fields{
				Http: http_lib.NewServer(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Api{
				Http: tt.fields.Http,
			}
			if err := a.defineRouters(); (err != nil) != tt.wantErr {
				t.Errorf("defineRouters() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewApi(t *testing.T) {
	tests := []struct {
		name string
		want Interface
	}{
		{
			name: "NewApi",
			want: NewApi(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewApi(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewApi() = %v, want %v", got, tt.want)
			}
		})
	}
}
