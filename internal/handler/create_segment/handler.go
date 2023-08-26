package create_segment

import (
	"github.com/gin-gonic/gin"
	"github.com/golovpeter/avito-trainee-task-2023/internal/common"
	"github.com/golovpeter/avito-trainee-task-2023/internal/service/create_segment"
	"github.com/sirupsen/logrus"
	"net/http"
)

type handler struct {
	service create_segment.CreateSegmentService
	log     *logrus.Logger
}

func NewHandler(
	log *logrus.Logger,
	service create_segment.CreateSegmentService,
) *handler {
	return &handler{
		log:     log,
		service: service,
	}
}

func (h *handler) CreateSegment(c *gin.Context) {
	var in CreateSegmentIn

	if err := c.BindJSON(&in); err != nil {
		h.log.WithError(err).Error("error binding JSON")
		c.JSON(http.StatusBadRequest, common.ErrorOut{
			ErrorMessage: "error binding JSON",
		})
		return
	}

	isValid, errMessage, err := validateInParams(in.SegmentSlug)

	if err != nil {
		h.log.WithError(err).Error(err.Error())
		c.JSON(http.StatusInternalServerError, common.ErrorOut{
			ErrorMessage: err.Error(),
		})
		return
	}

	if !isValid {
		h.log.Warn(errMessage)
		c.JSON(http.StatusBadRequest, common.ErrorOut{
			ErrorMessage: errMessage,
		})
		return
	}

	err = h.service.CreateSegment(&create_segment.CreateSegmentData{
		SegmentSlug: in.SegmentSlug,
	})

	if err != nil {
		h.log.WithError(err).Error("error create segments")
		c.JSON(http.StatusInternalServerError, common.ErrorOut{
			ErrorMessage: "error create segments",
		})
		return
	}
}
