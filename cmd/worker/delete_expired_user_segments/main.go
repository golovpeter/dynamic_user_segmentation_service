package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/golovpeter/avito-trainee-task-2023/internal/common"
	"github.com/golovpeter/avito-trainee-task-2023/internal/config"
	"github.com/golovpeter/avito-trainee-task-2023/internal/repository/user_segments"
	"github.com/golovpeter/avito-trainee-task-2023/internal/service/delete_expired_user_segments"
)

var signalChan = make(chan os.Signal, 1)

const tickerInterval = time.Minute

func main() {
	logger := logrus.New()

	cfg, err := config.Parse()
	if err != nil {
		return
	}

	dbConn, err := common.CreateDbClient(cfg.Database)
	if err != nil {
		logger.WithError(err).Error("error to create database client")
		return
	}

	userSegmentsRepository := user_segments.NewRepository(dbConn)
	deleteExpiredUserSegmentsService := delete_expired_user_segments.NewService(userSegmentsRepository)

	ticker := time.NewTicker(tickerInterval)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	logger.Info("time segment worker started")

	for {
		select {
		case <-ticker.C:
			err = deleteExpiredUserSegmentsService.DeleteExpiredUserSegments()
			if err != nil {
				logger.WithError(err).Error("error delete expired user segments")
			}
		case sigCaught := <-signalChan:
			logger.Infof("signal catched %s", sigCaught)
			return
		}
	}
}
