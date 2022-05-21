package api

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"service/src/api/router/scrapper"
	"service/src/internal/http_lib"
)

type Interface interface {
	Init() (err error)
	defineRouters() (err error)
}

type Api struct {
	Http http_lib.Interface
}

func NewApi() Interface {
	var instance = Api{
		Http: http_lib.NewServer(),
	}
	return &instance
}

func (a *Api) Init() (err error) {
	log.Info().Msg("Init")

	if err = a.defineRouters(); err != nil {
		return
	}
	if err = a.Http.Init(); err != nil {
		return
	}
	return
}

func (a *Api) defineRouters() (err error) {
	http.Handle("/scrap", scrapper.NewScrapper())

	return
}
