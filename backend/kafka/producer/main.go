package main

import (
	"common/ckafka"
	"common/logger"

	"producer/router"
)

func main() {
	logger.InitLogger()
	ckafka.InitKafka()

	r := router.NewRouter()

	//r.Use(interceptor.Recovery())

	if err := r.Run(":8080"); err != nil {
		logger.Error("gin run error", err)
		return
	}
}
