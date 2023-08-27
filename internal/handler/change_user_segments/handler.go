package change_user_segments

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golovpeter/avito-trainee-task-2023/internal/common"
	"github.com/golovpeter/avito-trainee-task-2023/internal/service/change_user_segments"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service change_user_segments.ChangeUserSegmentsService
	log     *logrus.Logger
}

func NewHandler(
	service change_user_segments.ChangeUserSegmentsService,
	log *logrus.Logger,
) *handler {
	return &handler{
		service: service,
		log:     log,
	}
}

// Delete segment godoc
// @Summary      Change user segments
// @Description	 delete and create user segments by name
// @Tags         segments
// @Accept       json
// @Produce      json
// @Param request body ChangeUserSegmentsIn true "request"
// @Success 200
// @Failure 400 {object} common.ErrorOut
// @Failure 500 {object} common.ErrorOut
// @Router       /segment/changeForUser [post]
func (h *handler) ChangeUserSegments(c *gin.Context) {
	var in ChangeUserSegmentsIn

	if err := c.BindJSON(&in); err != nil {
		h.log.WithError(err).Warn("error binding JSON")
		c.JSON(http.StatusBadRequest, common.ErrorOut{
			ErrorMessage: "error binding JSON",
		})
		return
	}

	valid, errMessage := validateInParams(&in)
	if !valid {
		h.log.Warn(errMessage)
		c.JSON(http.StatusBadRequest, common.ErrorOut{
			ErrorMessage: errMessage,
		})
		return
	}

	err := h.service.ChangeUserSegments(&change_user_segments.ChangeUserSegmentsData{
		AddSegments:    in.AddSegments,
		DeleteSegments: in.DeleteSegments,
		UserID:         in.UserID,
		ExpiredAt:      in.ExpiredAt,
	})

	if err != nil {
		switch err := err.(type) {
		case change_user_segments.ErrorSegmentsNotFound:
			h.log.Warn(err.Error())
			c.JSON(http.StatusBadRequest, common.ErrorOut{
				ErrorMessage: err.Error(),
			})

		default:
			h.log.WithError(err).Error("error change segments")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error_message": "error change segments",
			})
		}

		return
	}
}
