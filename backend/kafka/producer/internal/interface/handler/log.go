package handler

import (
	"context"
	"net/http"

	"producer/internal/interface/oapi"
	"time"

	"common/ckafka"
	"common/logger"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

func CreateLog(c *gin.Context) {
	var req *oapi.Log
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("c.ShouldBindJSON error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	msg := kafka.Message{
		Key:   []byte("preset"),
		Value: []byte(req.Pattern),
		Time:  time.Now(),
	}

	if err := ckafka.KafkaWriter.WriteMessages(context.Background(), msg); err != nil {
		logger.Error("failed to write message: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to write message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
