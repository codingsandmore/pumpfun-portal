package trades

import (
	"github.com/codingsandmore/pumpfun/portal"
	"github.com/codingsandmore/pumpfun/portal/decoders"
	"github.com/rs/zerolog/log"
)

type TrackedTokenServer struct {
	Tokens []string
}
type PriceTracker struct {
	Client portal.WebSocketClient
	trades chan any
}

type TrackRequest struct {
	Method string   `json:"method"`
	Keys   []string `json:"keys"`
}

func NewPriceTracker() *PriceTracker {
	tracker := &PriceTracker{
		Client: portal.NewPairClient(),
		trades: make(chan any),
	}

	go func() {
		err := tracker.Client.Subscribe(tracker.trades, &decoders.TradeDecoder{}, nil)
		if err != nil {
			log.Error().Err(err).Msg("Error subscribing to trades")
		}
	}()

	return tracker
}

func (t *PriceTracker) TrackPair(pair *portal.NewPairResponse) error {
	t.Client.Send(TrackRequest{
		Method: "subscribeTokenTrade",
		Keys:   []string{pair.Mint},
	})
	return nil
}

func (t *PriceTracker) SubscribeToTrades(trades chan *portal.NewTradeResponse) {
	log.Info().Msg("subscribing to channel for new trades")
	for {
		select {
		case m := <-t.trades:
			switch trade := m.(type) {
			case *portal.NewTradeResponse:
				if trade.Signature != "" {
					log.Debug().Msgf("received new trade data: %+v", m)
					trades <- trade
				}
			}
		}
	}
}
