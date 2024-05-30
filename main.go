package main

import (
	"github.com/codingsandmore/pumpfun-portal/portal"
	"github.com/codingsandmore/pumpfun-portal/portal/server"
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

	discoverPair := func(p *portal.NewPairResponse) {
		log.Info().Any("pair", p).Msgf("discovered pair")
	}
	discoverTrade := func(p *portal.NewTradeResponse) {
		log.Info().Any("trade", p).Msg("discovered trade")
	}

	server.NewPortalServer().Discover(discoverPair, discoverTrade)
	//shutdown on signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigs
	log.Info().Msgf("received signal: %s", sig)
}
