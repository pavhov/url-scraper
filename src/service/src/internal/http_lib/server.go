package http_lib

import (
	"net/http"
	"service/src/internal/builtin_lib"
	"service/src/tools/config_lib"
)

type Interface interface {
	Init() (err error)
}

type Server struct {
}

func NewServer() Interface {
	var instance Server
	return &instance
}

func (s *Server) Init() (err error) {
	defer builtin_lib.Recovery()

	err = http.ListenAndServe(config_lib.Config.Get("app_http_port").(string), nil)

	return
}
