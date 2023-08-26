package get_user_segments

import (
	"github.com/gin-gonic/gin"
	"github.com/golovpeter/avito-trainee-task-2023/internal/common"
	"github.com/golovpeter/avito-trainee-task-2023/internal/service/get_user_segments"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type handler struct {
	log     *logrus.Logger
	service get_user_segments.GetUserSegmentsService
}

func NewHandler(
	log *logrus.Logger,
	service get_user_segments.GetUserSegmentsService,
) *handler {
	return &handler{
		log:     log,
		service: service,
	}
}

func (h *handler) GetUserSegments(c *gin.Context) {
	userIdParam := c.Param("user_id")

	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		h.log.WithError(err).Warn("invalid user id")
		c.JSON(http.StatusBadRequest, common.ErrorOut{
			ErrorMessage: "invalid user id",
		})
		return
	}

	userSegments, err := h.service.GetUserSegments(
		&get_user_segments.GetUserSegmentsData{
			UserId: userId,
		})

	if err != nil {
		h.log.WithError(err).Error("error get user segments")
		c.JSON(http.StatusInternalServerError, common.ErrorOut{
			ErrorMessage: "error get user segments",
		})
	}

	c.JSON(http.StatusOK, GetUserSegmentsOut{
		Segments: userSegments,
	})
}
