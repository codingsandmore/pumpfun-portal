package decoders

import (
	"encoding/json"
	"github.com/codingsandmore/pumpfun-portal/portal"
	"github.com/rs/zerolog/log"
)

type TradeDecoder struct {
}

func (d *TradeDecoder) Decode(bytes []byte) (interface{}, error) {

	var result portal.NewTradeResponse

	err := json.Unmarshal(bytes, &result)

	if err != nil {
		log.Info().Msgf("Error decoding json: %v. Ignored message and moving on", err)
		return nil, err
	} else {
		return &result, nil
	}
}
