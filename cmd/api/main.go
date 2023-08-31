package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golovpeter/avito-trainee-task-2023/internal/cache/percent_segments"
	"github.com/golovpeter/avito-trainee-task-2023/internal/service/get_percent_segments"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"sync"
	"time"

	_ "github.com/golovpeter/avito-trainee-task-2023/docs"
	"github.com/golovpeter/avito-trainee-task-2023/internal/common"
	"github.com/golovpeter/avito-trainee-task-2023/internal/config"
	"github.com/golovpeter/avito-trainee-task-2023/internal/handler/change_user_segments"
	"github.com/golovpeter/avito-trainee-task-2023/internal/handler/create_segment"
	"github.com/golovpeter/avito-trainee-task-2023/internal/handler/delete_segment"
	"github.com/golovpeter/avito-trainee-task-2023/internal/handler/get_user_segments"
	"github.com/golovpeter/avito-trainee-task-2023/internal/repository/segments"
	"github.com/golovpeter/avito-trainee-task-2023/internal/repository/user_segments"
	change_user_segments_service "github.com/golovpeter/avito-trainee-task-2023/internal/service/change_user_segments"
	create_segment_service "github.com/golovpeter/avito-trainee-task-2023/internal/service/create_segment"
	delete_segment_service "github.com/golovpeter/avito-trainee-task-2023/internal/service/delete_segment"
	get_user_segments_service "github.com/golovpeter/avito-trainee-task-2023/internal/service/get_user_segments"
)

const percentSegmentsCacheUpdateInterval = time.Minute

// @title           Dynamic User Segmentation api Swagger API
// @version         1.0
// @description     API for Golang Dynamic User Segmentation api.
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8080
// @BasePath  /v1
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
		logger.WithError(err).Error("error to create database client")
		return
	}

	percentSegmentsCache := percent_segments.NewCache()

	segmentsRepository := segments.NewRepository(dbConn)
	userSegmentsRepository := user_segments.NewRepository(dbConn)

	changeUserSegmentsService := change_user_segments_service.NewService(segmentsRepository, userSegmentsRepository)
	createSegmentService := create_segment_service.NewService(segmentsRepository)
	deleteSegmentService := delete_segment_service.NewService(segmentsRepository)
	getUserSegmentsService := get_user_segments_service.NewService(userSegmentsRepository)
	getPercentService := get_percent_segments.NewService(segmentsRepository)

	createSegmentHandler := create_segment.NewHandler(logger, createSegmentService)
	deleteSegmentHandler := delete_segment.NewHandler(logger, deleteSegmentService, percentSegmentsCache)
	changeUserSegmentsHandler := change_user_segments.NewHandler(logger, changeUserSegmentsService)
	getUserSegmentsHandler := get_user_segments.NewHandler(logger, getUserSegmentsService, percentSegmentsCache)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go updatePercentSegmentsCache(getPercentService, percentSegmentsCache, &wg, logger)

	wg.Wait()

	router := gin.Default()

	router.POST("/v1/segment/create", createSegmentHandler.CreateSegment)
	router.POST("/v1/segment/delete", deleteSegmentHandler.DeleteSegment)
	router.POST("/v1/segment/changeForUser", changeUserSegmentsHandler.ChangeUserSegments)
	router.GET("/v1/segments/user/:user_id", getUserSegmentsHandler.GetUserSegments)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if err = router.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		logger.WithError(err).Error("server error occurred")
	}
}

func updatePercentSegmentsCache(
	getPercentService get_percent_segments.GetPercentSegmentsService,
	percentSegmentsCache *percent_segments.Cache,
	wg *sync.WaitGroup,
	logger *logrus.Logger,
) {
	// TODO избавиться от дублирующего кода
	percentSegments, err := getPercentService.GetPercentSegments()
	if err != nil {
		logger.WithError(err).Error("error to get percent segments")
	}

	percentSegmentsCache.Update(percentSegments)
	wg.Done()

	ticker := time.NewTicker(percentSegmentsCacheUpdateInterval)

	for {
		select {
		case <-ticker.C:
			percentSegments, err = getPercentService.GetPercentSegments()
			if err != nil {
				logger.WithError(err).Error("error to get percent segments")
			}

			percentSegmentsCache.Update(percentSegments)
		}
	}
}
