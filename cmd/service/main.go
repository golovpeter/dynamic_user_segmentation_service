package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golovpeter/avito-trainee-task-2023/internal/config"
	"github.com/golovpeter/avito-trainee-task-2023/internal/handler/create_segment"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
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

	db, err := sqlx.Connect("pgx",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			cfg.Database.Username,
			cfg.Database.Password,
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.Database))

	if err != nil {
		logger.WithError(err).Error("database connection error")
		return
	}

	if err = db.Ping(); err != nil {
		logger.WithError(err).Error("database access error")
		return
	}

	router := gin.Default()
	router.POST("api/v1/segment/create", create_segment.NewHandler(logger, db).CreateSegment)

	if err = router.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		logger.WithError(err).Error("server error occurred")
	}
}
