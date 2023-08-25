package create_segment

import (
	"errors"
	"regexp"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/golovpeter/avito-trainee-task-2023/internal/dto/create_segment"
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
		conn: conn,
		log:  log,
	}
}

func (h *handler) CreateSegment(c *gin.Context) {
	var in create_segment.CreateSegmentIn

	if err := c.BindJSON(&in); err != nil {
		h.log.WithError(err).Error("error binding JSON")
		c.JSON(400, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	err := validateSlug(in.SegmentSlug)
	if err != nil {
		h.log.Error(err)
		c.JSON(400, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	conn := segments.NewDbSegment(h.conn)

	err = conn.CreateSegment(in.SegmentSlug)
	if err != nil {
		h.log.WithError(err).Error("error adding a segment")
		c.JSON(400, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	c.Status(200)
}

func validateSlug(slug string) error {
	if utf8.RuneCountInString(slug) > 256 {
		return errors.New("line length exceeded")
	}

	validPattern := `^AVITO_[A-Z0-9_]+(_[A-Z0-9_]+)*$`
	isValid := regexp.MustCompile(validPattern).MatchString(slug)

	if !isValid {
		return errors.New("invalid slug format")
	}

	return nil
}
