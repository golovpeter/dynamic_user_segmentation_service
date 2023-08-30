package get_percent_segments

import "github.com/golovpeter/avito-trainee-task-2023/internal/repository/segments"

type GetPercentSegmentsService interface {
	GetPercentSegments() (map[string]segments.Segment, error)
}
