package system

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"reflect"
	mock_system "service/src/mock/init/system"
	"testing"
)

func TestNewSystem(t *testing.T) {
	tests := []struct {
		name string
		want Interface
	}{
		{
			name: "",
			want: NewSystem(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSystem(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSystem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystem_Init(t *testing.T) {
	tests := []struct {
		name string
		run  func(t *testing.T)
	}{
		{
			name: "System_Init",
			run: func(t *testing.T) {
				assertions := assert.New(t)
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				ms := mock_system.NewMockInterface(ctrl)
				ms.EXPECT().Init()

				if err := ms.Init(); err != nil {
					assertions.Fail(err.Error())
				}

				si := NewSystem()

				if err := si.Init(); err != nil {
					assertions.Fail(err.Error())
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.run(t)
		})
	}
}
