package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golovpeter/avito-trainee-task-2023/internal/config"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	cfg, err := config.Parse()
	if err != nil {
		logger.Error("error to parse config file")
		return
	}

	level, err := logrus.ParseLevel(cfg.Logger.Level)
	if err != nil {
		logger.Error("error to parse logger level")
		return
	}

	logger.SetLevel(level)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	if err = router.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		logger.WithError(err).Error("server error occurred")
	}
}
