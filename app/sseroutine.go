package app

import "github.com/mkamadeus/iot-smart-retail/handler/api"

type Routine interface {
	Run(done <-chan interface{})
}

type SSERoutine struct {
	ApiHandler *api.Handler
}

func NewRoutine(handlers *api.Handler) *SSERoutine {
	return &SSERoutine{
		ApiHandler: handlers,
	}
}

func (s *SSERoutine) Run(done <-chan interface{}) {
	go s.ApiHandler.SSE.Service.Broker(done)
}
