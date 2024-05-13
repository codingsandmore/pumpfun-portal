package portal

import "fmt"

type NewPairResponse struct {
	Signature             string  `json:"signature"`
	Mint                  string  `json:"mint"`
	TraderPublicKey       string  `json:"traderPublicKey"`
	TxType                string  `json:"txType"`
	InitialBuy            float64 `json:"initialBuy"`
	BondingCurveKey       string  `json:"bondingCurveKey"`
	VTokensInBondingCurve float64 `json:"vTokensInBondingCurve"`
	VSolInBondingCurve    float64 `json:"vSolInBondingCurve"`
	MarketCapSol          float64 `json:"marketCapSol"`
}

func (p *NewPairResponse) String() string {
	return fmt.Sprintf(
		"NewPairResponse: \n\tSignature: %s\n\tMint: %s\n\tTraderPublicKey: %s\n\tTxType: %s\n\tInitialBuy: %f\n\tBondingCurveKey: %s\n\tVTokensInBondingCurve: %f\n\tVSolInBondingCurve: %f\n\tMarketCapSol: %f",
		p.Signature,
		p.Mint,
		p.TraderPublicKey,
		p.TxType,
		p.InitialBuy,
		p.BondingCurveKey,
		p.VTokensInBondingCurve,
		p.VSolInBondingCurve,
		p.MarketCapSol,
	)
}

type NewTradeResponse struct {
	Signature             string  `json:"signature"`
	Mint                  string  `json:"mint"`
	TraderPublicKey       string  `json:"traderPublicKey"`
	TxType                string  `json:"txType"`
	TokenAmount           float64 `json:"tokenAmount"`
	NewTokenBalance       float64 `json:"newTokenBalance"`
	BondingCurveKey       string  `json:"bondingCurveKey"`
	VTokensInBondingCurve float64 `json:"vTokensInBondingCurve"`
	VSolInBondingCurve    float64 `json:"vSolInBondingCurve"`
	MarketCapSol          float64 `json:"marketCapSol"`
}

func (n *NewTradeResponse) String() string {
	return fmt.Sprintf(
		"NewTradeResponse: \n\tSignature: %s\n\tMint: %s\n\tTraderPublicKey: %s\n\tTxType: %s\n\tTokenAmount: %d\n\tNewTokenBalance: %f\n\tBondingCurveKey: %s\n\tVTokensInBondingCurve: %f\n\tVSolInBondingCurve: %f\n\tMarketCapSol: %f",
		n.Signature,
		n.Mint,
		n.TraderPublicKey,
		n.TxType,
		n.TokenAmount,
		n.NewTokenBalance,
		n.BondingCurveKey,
		n.VTokensInBondingCurve,
		n.VSolInBondingCurve,
		n.MarketCapSol,
	)
}
