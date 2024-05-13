package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"pumpfun/portal"
	"syscall"
	"time"
)

type SuccessResponse struct {
	Message string `json:"message"`
}

func main() {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	//log.SetReportCaller(true)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	//subscribe to web socket

	client := portal.NewPairClient()

	var decoder portal.MessageDecoder
	pairs := make(chan any)
	go func() { client.Subscribe(pairs, decoder) }()

	time.Sleep(5 * time.Second)
	client.Send("{\"method\" : \"subscribeNewToken\"}")

	for {
		select {
		case m := <-pairs:
			log.Info().Any("message", m).Msg("Received message")
		}
	}
	//shutdown on signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigs
	log.Info().Msgf("received signal: %s", sig)
}
