package main

import (
	"github.com/codingsandmore/pumpfun/portal"
	"github.com/codingsandmore/pumpfun/portal/decoders"
	"github.com/codingsandmore/pumpfun/portal/prices"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	//log.SetReportCaller(true)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	//subscribe to web socket

	trackNewPairs()
	//shutdown on signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigs
	log.Info().Msgf("received signal: %s", sig)
}

func trackNewPairs() {
	client := portal.NewPairClient()

	pairs := make(chan any)
	trackedTrades := make(chan *portal.NewTradeResponse)
	tracker := prices.NewPriceTracker()

	go func() {

		err := client.Subscribe(pairs, &decoders.PairDecoder{}, "{\"method\" : \"subscribeNewToken\"}")

		if err != nil {
			log.Fatal().Err(err).Msgf("failed to subscribe to pairs. Error was %v", err)
		}
	}()

	go func() {
		for {
			select {
			case m := <-pairs:
				go func(p *portal.NewPairResponse) {
					if p == nil || p.Signature == "" {
						log.Debug().Msgf("failed to subscribe to pair, due to being Nil or none populated")
						return
					}
					err := tracker.TrackPair(p)
					if err != nil {
						log.Error().Err(err).Msg("failed to track pair")
					}
				}(m.(*portal.NewPairResponse))
			}
		}
	}()

	go tracker.SubscribeToTrades(trackedTrades)

	for {
		select {
		case m := <-trackedTrades:
			log.Info().Any("trade", m).Msgf("tracking trade")
		}
	}
}
