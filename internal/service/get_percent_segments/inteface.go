package get_percent_segments

import (
	"github.com/golovpeter/avito-trainee-task-2023/internal/cache/percent_segments"
)

type GetPercentSegmentsService interface {
	GetPercentSegments() (map[string]percent_segments.Segment, error)
}
