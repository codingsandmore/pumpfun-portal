package exchange

type Exchange interface {
	SwapSolForToken(amountSol float64, token string, slippage float64, priorityFee float64) (any, error)

	SwapTokenForSol(amountToken float64, token string, slippageToken float64, priorityFee float64) (any, error)
}
