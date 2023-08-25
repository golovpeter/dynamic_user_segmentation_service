package delete_segment

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golovpeter/avito-trainee-task-2023/internal/dto/delete_segment_dto"
	"github.com/golovpeter/avito-trainee-task-2023/internal/repository/segments"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type handler struct {
	log  *logrus.Logger
	conn *sqlx.DB
}

func NewHandler(log *logrus.Logger, conn *sqlx.DB) *handler {
	return &handler{
		log:  log,
		conn: conn,
	}
}

func (h *handler) DeleteSegment(c *gin.Context) {
	var in delete_segment_dto.DeleteSegmentIn

	if err := c.BindJSON(&in); err != nil {
		h.log.WithError(err).Error("error binding JSON")
		c.JSON(400, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	conn := segments.NewDbSegment(h.conn)

	err, result := conn.DeleteSegment(in.SegmentSlug)
	if err != nil {
		h.log.WithError(err).Error("error deleting a segment")
		c.JSON(400, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		err := errors.New("this segment was not found")

		h.log.Error(err)
		c.JSON(400, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	c.Status(200)
}
