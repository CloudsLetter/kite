package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"kite/middlewares"
	"kite/router"
	"kite/utilities"
	"log"
	"sync"
)

var PNGProcessingChannel = make(chan string)

func main() {
	//numCPU := runtime.NumCPU()

	//maxConcurrent := numCPU / 2 // 设置并发量为 CPU 核心数的 2 倍

	var once sync.Once
	//
	//var wg sync.WaitGroup
	//counterCh := make(chan int)

	gin.SetMode(gin.ReleaseMode)
	once.Do(utilities.PostgresInit)

	r := gin.Default()
	r.Use(middlewares.Cors())
	r.MaxMultipartMemory = 8 << 20
	/*	r.Get("/images", "./data/Images")
	 */
	r.POST("/uploadimage", router.UpLoadImage)
	r.POST("/uploadurlimage", router.UploadUrImg)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Gin Server Run Error: %v", err)
		return
	}
}
