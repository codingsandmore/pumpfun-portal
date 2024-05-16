package exchange

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

type PumpFun struct {
	url    string
	apiKey string
	pool   string
	client *http.Client
}

type SwapRequest struct {
	Action           string  `json:"action"`
	Mint             string  `json:"mint"`
	Amount           float64 `json:"amount"`
	DenominatedInSol bool    `json:"denominatedInSol,string"`
	Slippage         float64 `json:"slippage"`
	PriorityFee      float64 `json:"priorityFee"`
	Pool             string  `json:"pool"`
}

type SwapResponse struct {
	Errors []string `json:"errors"`
}

func (p *PumpFun) SwapSolForToken(amountSol float64, token string, slippage float64, priorityFee float64) (any, error) {

	request := SwapRequest{
		Action:           "buy",
		Mint:             token,
		Amount:           amountSol,
		DenominatedInSol: true,
		Slippage:         slippage,
		PriorityFee:      priorityFee,
		Pool:             p.pool,
	}

	return p.SwapRequest(request)
}

func (p *PumpFun) SwapRequest(request SwapRequest) (any, error) {
	/*
		if request.Amount <= 0 {
			return nil, errors.New("invalid swap request, need to specify an amount larger than 0")
		}
		if request.Mint == "" {
			return nil, errors.New("invalid swap request, need to specify a mint address")
		}
		if request.Action == "buy" || request.Action == "sell" {

		} else {
			return nil, errors.New("invalid swap request, need to specify a valid action")
		}
		if request.Slippage <= 0 {
			return nil, errors.New("invalid swap request, need to specify a slippage larger than 0")
		}
		if request.PriorityFee <= 0 {
			return nil, errors.New("invalid swap request, need to specify a priority fee larger than 0")
		}
		if p.apiKey == "" {
			return nil, errors.New("invalid swap request, need to specify an api key")
		}
	*/
	api := fmt.Sprintf("%s?api-key=%s", p.url, p.apiKey)

	jsonData, err := json.Marshal(request)

	if err != nil {
		log.Error().Err(err).Msg("Error marshalling request")
		return nil, err
	}

	resp, err := p.client.Post(api, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error().Err(err).Msg("Error posting request")
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		bodyString := string(bodyBytes)

		log.Info().Str("state", resp.Status).Str("content", bodyString).Msg("Request response")

		return nil, nil
	}
}

func (p *PumpFun) SwapTokenForSol(amountToken float64, token string, slippage float64, priorityFee float64) (any, error) {

	request := SwapRequest{
		Action:           "sell",
		Mint:             token,
		Amount:           amountToken,
		DenominatedInSol: false,
		Slippage:         slippage,
		PriorityFee:      priorityFee,
		Pool:             p.pool,
	}

	return p.SwapRequest(request)
}

func NewPumpFun(apiKey string, client *http.Client) *PumpFun {
	return &PumpFun{
		pool:   "pump",
		apiKey: apiKey,
		url:    "https://pumpportal.fun/api/trade",
		client: client,
	}
}
