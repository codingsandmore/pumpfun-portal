package exchange

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/rs/zerolog/log"
	"golang.org/x/time/rate"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"time"
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
	Signature string `json:"signature"`
	Errors    []any  `json:"errors"`
	Request   SwapRequest
	Filled    float64
}

func (p *PumpFun) SwapSolForToken(amountSol float64, token string, slippage float64, priorityFee float64) (*SwapResponse, error) {

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

func (p *PumpFun) SwapRequest(request SwapRequest) (*SwapResponse, error) {
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
		sr := SwapResponse{}
		body, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			log.Error().Err(readErr).Msgf("Error reading response")
		}

		log.Info().Str("content", string(body)).Msg("Response body")
		jsonErr := json.Unmarshal(body, &sr)
		if jsonErr != nil {
			log.Error().Err(jsonErr).Msgf("Error unmarshalling response")
		}

		log.Info().Str("state", resp.Status).Any("response", sr).Msg("Request response")

		sr.Request = request
		if len(sr.Errors) > 0 {
			return &sr, errors.New("we received an error during this response")
		}

		filled, err := p.ParseTransaction(&sr)

		if err != nil {
			return &sr, err
		}
		sr.Filled = filled
		return &sr, nil
	}
}

func (p *PumpFun) ParseTransaction(sr *SwapResponse) (float64, error) {
	client := rpc.NewWithCustomRPCClient(rpc.NewWithLimiter(
		"https://alpha-quick-energy.solana-mainnet.quiknode.pro/a9a59fa0f479c337641331fadd97fb1239fd4744/",
		rate.Every(time.Second), // time frame
		5,                       // limit of requests per time frame
	))

	var v uint64 = 1

	opts := rpc.GetParsedTransactionOpts{
		MaxSupportedTransactionVersion: &v,
	}
	sig, err := solana.SignatureFromBase58(sr.Signature)
	if err != nil {
		return 0, err
	}

	var t *rpc.GetParsedTransactionResult
	for i := 0; ; i++ {
		t, err = client.GetParsedTransaction(context.TODO(), sig, &opts)

		if err == nil {
			//done
			log.Info().Any("tx", t)
			break
		} else if err.Error() == "not found" {
			log.Debug().AnErr("error", err).Msgf("received error, sleeping a bit...")
			time.Sleep(1 * time.Second)
		} else {
			log.Error().AnErr("error", err)
			return 0, err
		}

	}
	log.Info().Any("tx", t).Msgf("fetched transactions")

	if t.Meta.Err != nil {
		log.Error().Str("pair", sr.Signature).Any("error in transaction", t.Meta.Err).Msgf("transaction most likely failed")
		return 0, err
	}

	if sr.Request.Action == "buy" {
		tb := p.findTokens(sr.Request.Mint, t.Meta.PreTokenBalances)
		ta := p.findTokens(sr.Request.Mint, t.Meta.PostTokenBalances)
		tokensBefore, err1 := strconv.ParseFloat(tb.Amount, 64)
		tokensAfter, err2 := strconv.ParseFloat(ta.Amount, 64)

		if err1 != nil {
			return 0, err1
		}
		if err2 != nil {
			return 0, err2
		}
		difference := tokensBefore - tokensAfter
		adjusted := difference / math.Pow(10, float64(tb.Decimals))

		log.Info().Float64("gained", difference).Msgf("we received an amount")

		return adjusted, nil
	} else {
		solBefore := float64(t.Meta.PreBalances[0]) / 1000000000
		solAfter := float64(t.Meta.PostBalances[0]) / 1000000000
		gain := solAfter - solBefore

		return gain, nil

	}
}

func (p *PumpFun) findTokens(mint string, balance []rpc.TokenBalance) *rpc.UiTokenAmount {
	for _, token := range balance {
		if token.Mint.String() == mint {
			return token.UiTokenAmount
		}
	}

	return nil
}

func (p *PumpFun) SwapTokenForSol(amountToken float64, token string, slippage float64, priorityFee float64) (*SwapResponse, error) {

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
