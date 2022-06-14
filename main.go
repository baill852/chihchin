package main

import (
	"encoding/json"
	"test/lib"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config := lib.LoadConfig()
	wsClient := lib.NewWss(config)
	redisClient := lib.NewRedis(config)

	go func() {
		for {
			data, _ := wsClient.SubscribeRecentDealsByStream("btcusdt@aggTrade")
			b, _ := json.Marshal(data)
			redisClient.Set("streams=btcusdt@aggTrade", string(b), time.Duration(time.Hour*24))
			time.Sleep(time.Second * 1)
		}
	}()

	server := gin.Default()
	server.GET("/", func(ctx *gin.Context) {
		c := redisClient.Get("streams=btcusdt@aggTrade")
		if s, err := c.Result(); err != nil {
			ctx.JSON(400, err)
		} else {
			data := map[string]interface{}{}
			json.Unmarshal([]byte(s), &data)
			ctx.JSON(200, data)
		}
	})
	server.Run("0.0.0.0:8888")
}
