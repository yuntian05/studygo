package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

const keyRequestId = "requestId"

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		// log path  latency status
		logger.Info("incomming request",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(start)))
	}, func(c *gin.Context) {
		c.Set(keyRequestId, rand.Int())
		c.Next()
	})
	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, existed := c.Get(keyRequestId); existed {
			h[keyRequestId] = rid
		}
		c.JSON(200, h)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	err = r.Run()
	if err != nil {
		panic(err)
	}
}
