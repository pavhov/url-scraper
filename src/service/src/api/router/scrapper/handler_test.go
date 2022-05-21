package scrapper

import (
	"net/http"
	"reflect"
	"service/src/internal/builtin_lib"
	"testing"
	"time"
)

func TestNewScrapper(t *testing.T) {
	tests := []struct {
		name string
		want Interface
	}{
		{
			name: "NewScrapper",
			want: NewScrapper(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScrapper(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScrapper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScrapper_ServeHTTP(t *testing.T) {
	type fields struct {
		limiter *builtin_lib.Limiter
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Scrapper_ServeHTTP",
			fields: fields{
				limiter: builtin_lib.NewLimiter(builtin_lib.Every(time.Second), 100),
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scrapper{
				limiter: tt.fields.limiter,
			}
			s.ServeHTTP(tt.args.w, tt.args.r)
		})
	}
}

func TestScrapper_checkLimit(t *testing.T) {
	type fields struct {
		limiter *builtin_lib.Limiter
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Scrapper_checkLimit",
			fields: fields{
				limiter: builtin_lib.NewLimiter(builtin_lib.Every(time.Second), 100),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scrapper{
				limiter: tt.fields.limiter,
			}
			if err := s.checkLimit(); (err != nil) != tt.wantErr {
				t.Errorf("checkLimit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestScrapper_checkMethod(t *testing.T) {
	type fields struct {
		limiter *builtin_lib.Limiter
	}
	type args struct {
		method string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Scrapper_checkMethod",
			fields: fields{
				limiter: builtin_lib.NewLimiter(builtin_lib.Every(time.Second), 100),
			},
			args: args{
				method: "POST",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scrapper{
				limiter: tt.fields.limiter,
			}
			if err := s.checkMethod(tt.args.method); (err != nil) != tt.wantErr {
				t.Errorf("checkMethod() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestScrapper_makeCalls(t *testing.T) {
	type fields struct {
		limiter *builtin_lib.Limiter
	}
	type args struct {
		url string
		ch  chan string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Scrapper_makeCalls",
			fields: fields{
				limiter: builtin_lib.NewLimiter(builtin_lib.Every(time.Second), 100),
			},
			args: args{
				url: "https://google.com",
				ch:  make(chan string, 1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scrapper{
				limiter: tt.fields.limiter,
			}
			s.makeCalls(tt.args.url, tt.args.ch)
		})
	}
}

func TestScrapper_parseData(t *testing.T) {
	type fields struct {
		limiter *builtin_lib.Limiter
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes []string
	}{
		{
			name: "Scrapper_parseData",
			fields: fields{
				limiter: builtin_lib.NewLimiter(builtin_lib.Every(time.Second), 100),
			},
			args: args{
				data: []byte(`
				https://google.com
				https://facebook.com
				`),
			},
			wantRes: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scrapper{
				limiter: tt.fields.limiter,
			}
			if gotRes := s.parseData(tt.args.data); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("parseData() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestScrapper_prepareCalls(t *testing.T) {
	type fields struct {
		limiter *builtin_lib.Limiter
	}
	type args struct {
		data []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes []byte
	}{
		{
			name: "Scrapper_prepareCalls",
			fields: fields{
				limiter: builtin_lib.NewLimiter(builtin_lib.Every(time.Second), 100),
			},
			args: args{
				data: []string{
					"https://google -- 1000",
					"https://facebook.com -- 1000",
				},
			},
			wantRes: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scrapper{
				limiter: tt.fields.limiter,
			}
			if gotRes := s.prepareCalls(tt.args.data); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("prepareCalls() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
