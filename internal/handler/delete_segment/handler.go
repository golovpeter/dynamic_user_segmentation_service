package delete_segment

import (
	"errors"
	"github.com/golovpeter/avito-trainee-task-2023/internal/common"
	"github.com/golovpeter/avito-trainee-task-2023/internal/service/delete_segment"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type handler struct {
	log     *logrus.Logger
	service delete_segment.DeleteSegmentService
}

func NewHandler(
	log *logrus.Logger,
	service delete_segment.DeleteSegmentService,
) *handler {
	return &handler{
		log:     log,
		service: service,
	}
}

func (h *handler) DeleteSegment(c *gin.Context) {
	var in DeleteSegmentIn

	if err := c.BindJSON(&in); err != nil {
		h.log.WithError(err).Error("error binding JSON")
		c.JSON(http.StatusBadRequest, common.ErrorOut{
			ErrorMessage: "error binding JSON",
		})
		return
	}

	err := h.service.DeleteSegmentService(&delete_segment.DeleteSegmentData{
		SegmentSlug: in.SegmentSlug,
	})

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

	if err != nil {
		switch {
		case errors.Is(err, delete_segment.ErrSegmentNotFound):
			h.log.WithError(err).Warn(err.Error())
			c.JSON(http.StatusBadRequest, common.ErrorOut{
				ErrorMessage: err.Error(),
			})
		default:
			h.log.WithError(err).Error("error deleting segment")
			c.JSON(http.StatusInternalServerError, common.ErrorOut{
				ErrorMessage: "error deleting segment",
			})
		}

		return
	}
}
