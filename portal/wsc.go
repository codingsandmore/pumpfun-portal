package portal

import (
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"net/url"
	"time"
)

type MessageDecoder interface {
	Decode(bytes []byte) (any, error)
}
type WebSocketClient interface {
	Connect() error
	Send(any)
	Subscribe(messasge chan any, decoder MessageDecoder)
}
type DefaultWebSocketClient struct {
	Addr             string
	Conn             *websocket.Conn
	IncomingMessages chan any
	OutgoingMessages chan any
	Decoder          MessageDecoder
}

func NewDefaultClient(addr string) *DefaultWebSocketClient {

	return &DefaultWebSocketClient{
		Addr:             addr,
		Conn:             nil,
		IncomingMessages: make(chan any),
		OutgoingMessages: make(chan any, 10), // Change 10 to appropriate buffer size
	}
}

func (c *DefaultWebSocketClient) Subscribe(messages chan any, decoder MessageDecoder) {
	c.Decoder = decoder
	c.Start()
	for {
		select {
		case m := <-c.IncomingMessages:
			log.Info().Msgf("message: %v", m)
			messages <- m
		}
	}
}
func (c *DefaultWebSocketClient) Connect() error {
	log.Info().Str("url", c.Addr).Msgf("connecting to %v", c.Addr)
	u, err := url.Parse(c.Addr)
	if err != nil {
		return err
	}

	c.Conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *DefaultWebSocketClient) maintainConnection() {
	for {
		if c.Conn == nil {
			if err := c.Connect(); err != nil {
				log.Info().Msgf("Connect Error:", err)
				time.Sleep(time.Second)
				continue
			}
		}

		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Error().Msgf("Read Error:", err)
			c.Conn = nil
			continue
		}

		message, err := c.Decoder.Decode(msg)

		if err != nil {
			log.Error().Msgf("Decode Error: %s", err)
		} else if message == nil {
			log.Warn().Msgf("ignored nil message")
		} else {
			log.Debug().Any("message", message).Msgf("forwarding decoded message")
			c.IncomingMessages <- message
		}
	}
}

func (c *DefaultWebSocketClient) maintainOutgoingMessages() {
	for msg := range c.OutgoingMessages {
		if c.Conn == nil {
			log.Error().Msgf("Connection Error: No active connection")
			continue
		}

		var err error
		switch v := msg.(type) {
		case []byte:
			err = c.Conn.WriteMessage(1, v)
		case string:
			err = c.Conn.WriteMessage(1, []byte(v))

		default:
			err = c.Conn.WriteJSON(v)
		}

		if err != nil {
			log.Error().Msgf("Write Error:", err)
			c.Conn = nil
			continue
		}
	}
}

func (c *DefaultWebSocketClient) Send(msg any) {
	log.Debug().Any("message", msg).Msg("Sending message")
	c.OutgoingMessages <- msg
}

func (c *DefaultWebSocketClient) Start() {
	go c.maintainConnection()
	go c.maintainOutgoingMessages()
}
