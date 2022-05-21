package scrapper

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"service/src/internal/builtin_lib"
	"strings"
	"time"
)

type (
	Interface interface {
		ServeHTTP(w http.ResponseWriter, r *http.Request)
		checkMethod(method string) (err error)
	}
	Scrapper struct {
		limiter *builtin_lib.Limiter
	}
)

func NewScrapper() Interface {
	var instance Scrapper
	instance.limiter = builtin_lib.NewLimiter(builtin_lib.Every(time.Second), 100)
	log.Info().Msg("Scrapper")
	return &instance
}

func (s *Scrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer builtin_lib.Recovery()

	if err := s.checkMethod(r.Method); err != nil {
		w.WriteHeader(405)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			log.Error().Msg(err.Error())
			return
		}
		return
	}
	if err := s.checkLimit(); err != nil {
		w.WriteHeader(429)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			log.Error().Msg(err.Error())
			return
		}
		return
	}

	if body, err := io.ReadAll(r.Body); err != nil {
		log.Error().Msg(err.Error())
		return
	} else {
		data := s.parseData(body)
		res := s.prepareCalls(data)

		w.WriteHeader(201)
		if _, err := w.Write(res); err != nil {
			log.Error().Msg(err.Error())
		}
	}
}

func (s *Scrapper) checkMethod(method string) (err error) {
	if method != "POST" {
		err = errors.New("wrong method")
	}
	return
}

func (s *Scrapper) checkLimit() (err error) {
	if !s.limiter.Allow() {
		err = errors.New("limit expected")
	}
	return
}

func (s *Scrapper) parseData(data []byte) (res []string) {
	swapSlice := strings.Split(string(data), "\n")
	for _, str := range swapSlice {
		if str != "" {
			res = append(res, str)
		}
	}
	return
}

func (s *Scrapper) prepareCalls(data []string) (res []byte) {
	if len(data) == 0 {
		return []byte("")
	}
	cMsg := make(chan string, len(data))

	for _, url := range data {
		go s.makeCalls(url, cMsg)
	}

	var strRes []string
	var state int
	for swap := range cMsg {
		strRes = append(strRes, swap)
		state++
		if state == len(data) {
			close(cMsg)
		}
	}
	return []byte(strings.Join(strRes, "\n"))
}

func (s *Scrapper) makeCalls(url string, ch chan string) {
	swap := strings.Replace(url, "\r", "", 1)
	if resp, err := http.Get(swap); err != nil {
		log.Error().Msg(err.Error())
		ch <- ""
	} else {
		if body, err := io.ReadAll(resp.Body); err != nil {
			log.Error().Msg(err.Error())
			return
		} else {
			ch <- fmt.Sprintf("%v - %v", swap, len(body))
		}
	}
}
