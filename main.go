package main

import (
	"github.com/codingsandmore/pumpfun/portal"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
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
	go func() {

		err := client.Subscribe(pairs, decoder)

		if err != nil {
			log.Fatal().Err(err).Msg("failed to subscribe to pairs")
		}
	}()

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
