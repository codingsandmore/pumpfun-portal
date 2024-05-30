package server

import (
	"github.com/codingsandmore/pumpfun-portal/portal"
	"github.com/codingsandmore/pumpfun-portal/portal/decoders"
	"github.com/codingsandmore/pumpfun-portal/portal/trades"
	"github.com/rs/zerolog/log"
)

type NewPairDiscovered func(p *portal.NewPairResponse)
type NewTradeDiscovered func(p *portal.NewTradeResponse)

type Server interface {
	Discover(pairDiscovery NewPairDiscovered, tradeDiscovery NewTradeDiscovered)
}

type PortalServer struct {
}

func NewPortalServer() *PortalServer {
	return &PortalServer{}
}

// Discover is a method of the PortalServer type that allows for the discovery of new pairs and trades.
// It takes two parameters: pairDiscovery of type NewPairDiscovered and tradeDiscovery of type NewTradeDiscovered.
// It creates a new pair client, initializes the pairs and trackedTrades channels, and creates a new trade tracker.
// It starts two goroutines: one to subscribe to pairs and one to track trades.
// The method then enters a loop to handle incoming trade information.
func (s *PortalServer) Discover(pairDiscovery NewPairDiscovered, tradeDiscovery NewTradeDiscovered) {
	client := portal.NewPairClient()
	defer client.Shutdown()

	pairs := make(chan any)
	defer close(pairs)

	trackedTrades := make(chan *portal.NewTradeResponse)
	defer close(trackedTrades)

	tracker := trades.NewPriceTracker()

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
					go pairDiscovery(p)
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
			tradeDiscovery(m)
		}
	}
}
