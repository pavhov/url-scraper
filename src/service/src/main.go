package main

import (
	"github.com/rs/zerolog/log"
	"os"
	"service/src/api"
	"service/src/init/system"
)

func init() {

}

func main() {
	if err := system.NewSystem().Init(); err != nil {
		log.Fatal().Msg(err.Error())
		os.Exit(1)
	}
	if err := api.NewApi().Init(); err != nil {
		log.Fatal().Msg(err.Error())
		os.Exit(1)
	}
}
