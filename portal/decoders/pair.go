package decoders

import (
	"encoding/json"
	"github.com/codingsandmore/pumpfun-portal/portal"
	"github.com/rs/zerolog/log"
)

type PairDecoder struct {
}

func (d *PairDecoder) Decode(bytes []byte) (interface{}, error) {

	var result portal.NewPairResponse

	err := json.Unmarshal(bytes, &result)

	if err != nil {
		log.Info().Msgf("Error decoding json: %v. Ignored message and moving on", err)
		return nil, err
	} else {
		return &result, nil
	}
}
