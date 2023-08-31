package delete_segment

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/golovpeter/avito-trainee-task-2023/internal/cache/percent_segments"
	"github.com/golovpeter/avito-trainee-task-2023/internal/common"
	"github.com/golovpeter/avito-trainee-task-2023/internal/service/delete_segment"
)

type handler struct {
	log     *logrus.Logger
	service delete_segment.DeleteSegmentService
	cache   *percent_segments.Cache
}

func NewHandler(
	log *logrus.Logger,
	service delete_segment.DeleteSegmentService,
	cache *percent_segments.Cache,
) *handler {
	return &handler{
		log:     log,
		service: service,
		cache:   cache,
	}
}

// Delete segment godoc
// @Summary      Delete segment
// @Description	 deleting segment by name
// @Tags         segments
// @Accept       json
// @Produce      json
// @Param request body DeleteSegmentIn true "request"
// @Success 200
// @Failure 400 {object} common.ErrorOut
// @Failure 500 {object} common.ErrorOut
// @Router       /segment/delete [post]
func (h *handler) DeleteSegment(c *gin.Context) {
	var in DeleteSegmentIn

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
			ErrorMessage: "failed parse regexp",
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

	err = h.service.DeleteSegment(&delete_segment.DeleteSegmentData{
		SegmentSlug:          in.SegmentSlug,
		PercentSegmentsCache: h.cache,
	})

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
