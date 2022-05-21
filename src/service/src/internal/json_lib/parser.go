package json_lib

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
)

func Encode[I any](data I) string {
	b, err := json.Marshal(data)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
	return string(b)
}

func Decode[I any](data string) I {
	var i I
	err := json.Unmarshal([]byte(data), &i)
	if err != nil {
		log.Error().Msgf(err.Error())
	}

	return i
}
