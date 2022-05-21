package config_lib

import (
	"reflect"
	"sync"
	"testing"
)

func TestConf_Get(t *testing.T) {
	type fields struct {
		Mutex  sync.Mutex
		config map[string]any
	}
	type args struct {
		k string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conf{
				Mutex:  tt.fields.Mutex,
				config: tt.fields.config,
			}
			if got := c.Get(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConf_Init(t *testing.T) {
	type fields struct {
		Mutex  sync.Mutex
		config map[string]any
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conf{
				Mutex:  tt.fields.Mutex,
				config: tt.fields.config,
			}
			if err := c.Init(); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConf_defineConfig(t *testing.T) {
	type fields struct {
		Mutex  sync.Mutex
		config map[string]any
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conf{
				Mutex:  tt.fields.Mutex,
				config: tt.fields.config,
			}
			if err := c.defineConfig(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("defineConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConf_prepareConfigs(t *testing.T) {
	type fields struct {
		Mutex  sync.Mutex
		config map[string]any
	}
	type args struct {
		data *map[string]any
		key  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conf{
				Mutex:  tt.fields.Mutex,
				config: tt.fields.config,
			}
			c.prepareConfigs(tt.args.data, tt.args.key)
		})
	}
}

func TestConf_setArr(t *testing.T) {
	type fields struct {
		Mutex  sync.Mutex
		config map[string]any
	}
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conf{
				Mutex:  tt.fields.Mutex,
				config: tt.fields.config,
			}
			c.setArr(tt.args.key, tt.args.val)
		})
	}
}

func TestConf_setBoll(t *testing.T) {
	type fields struct {
		Mutex  sync.Mutex
		config map[string]any
	}
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conf{
				Mutex:  tt.fields.Mutex,
				config: tt.fields.config,
			}
			c.setBoll(tt.args.key, tt.args.val)
		})
	}
}

func TestConf_setFloat(t *testing.T) {
	type fields struct {
		Mutex  sync.Mutex
		config map[string]any
	}
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conf{
				Mutex:  tt.fields.Mutex,
				config: tt.fields.config,
			}
			c.setFloat(tt.args.key, tt.args.val)
		})
	}
}

func TestNewConf(t *testing.T) {
	tests := []struct {
		name string
		want *Conf
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConf(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConf() = %v, want %v", got, tt.want)
			}
		})
	}
}
