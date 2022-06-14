package lib

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type Wss struct {
	Config Config
}

func NewWss(config *Config) *Wss {
	return &Wss{
		Config: *config,
	}
}

func (s *Wss) SubscribeRecentDealsByStream(stream string) (data map[string]interface{}, err error) {
	c, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("%s?streams=%s", s.Config.Ws.Host, stream), nil)
	if err != nil {
		return data, err
	}
	defer c.Close()

	_, b, err := c.ReadMessage()
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return data, err
	}

	if value, check := data["data"].(map[string]interface{}); check {
		return value, err
	}

	return data, err
}
