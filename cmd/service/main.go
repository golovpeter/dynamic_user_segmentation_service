package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/golovpeter/avito-trainee-task-2023/internal/common"
	"github.com/golovpeter/avito-trainee-task-2023/internal/config"
	"github.com/golovpeter/avito-trainee-task-2023/internal/handler/change_user_segments"
	"github.com/golovpeter/avito-trainee-task-2023/internal/handler/create_segment"
	"github.com/golovpeter/avito-trainee-task-2023/internal/handler/delete_segment"
	"github.com/golovpeter/avito-trainee-task-2023/internal/repository/segments"
	"github.com/golovpeter/avito-trainee-task-2023/internal/repository/user_segments"
	"github.com/golovpeter/avito-trainee-task-2023/internal/service/change_user_segments_service"
	create_segment_service "github.com/golovpeter/avito-trainee-task-2023/internal/service/create_segment"
	delete_segment_service "github.com/golovpeter/avito-trainee-task-2023/internal/service/delete_segment"
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

	dbConn, err := common.CreateDbClient(cfg.Database)
	if err != nil {
		logrus.WithError(err).Error("error to create database client")
		return
	}

	router := gin.Default()

	segmentsRepository := segments.NewRepository(dbConn)
	userSegmentsRepository := user_segments.NewRepository(dbConn)

	changeUserSegmentsService := change_user_segments_service.NewService(segmentsRepository, userSegmentsRepository)
	createSegmentService := create_segment_service.NewService(segmentsRepository)
	deleteSegmentService := delete_segment_service.NewService(segmentsRepository)

	createSegmentHandler := create_segment.NewHandler(logger, createSegmentService)
	deleteSegmentHandler := delete_segment.NewHandler(logger, deleteSegmentService)
	changeUserSegmentsHandler := change_user_segments.NewHandler(changeUserSegmentsService, logger)

	router.POST("/v1/segment/create", createSegmentHandler.CreateSegment)
	router.POST("/v1/segment/delete", deleteSegmentHandler.DeleteSegment)
	router.POST("/v1/segment/changeForUser", changeUserSegmentsHandler.ChangeUserSegments)

	if err = router.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		logger.WithError(err).Error("server error occurred")
	}
}
