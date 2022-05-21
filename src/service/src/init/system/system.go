package system

import (
	"github.com/rs/zerolog/log"
	"service/src/third_party/log_lib"
	"service/src/tools/config_lib"
)

type Interface interface {
	Init() (err error)
}

type System struct {
	Config config_lib.Interface
	Log    log_lib.Interface
}

func NewSystem() Interface {
	var instance = System{
		Config: config_lib.NewConf(),
		Log:    log_lib.NewLog(),
	}
	return &instance
}

func (i *System) Init() (err error) {
	log.Info().Msg("Init")
	if err = i.Config.Init(); err != nil {
		return
	}
	if err = i.Log.Init(); err != nil {
		return
	}
	return
}
