package portal

type PairClient struct {
	DefaultWebSocketClient
}

func NewPairClient() WebSocketClient {
	return &PairClient{*NewDefaultClient("wss://pumpportal.fun/api/data")}
}
