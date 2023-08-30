package get_percent_segments

//go:generate mockgen -destination=mocks.go -package=$GOPACKAGE -source=interfaces.go

import (
	"github.com/golovpeter/avito-trainee-task-2023/internal/cache/percent_segments"
)

type GetPercentSegmentsService interface {
	GetPercentSegments() (map[string]percent_segments.Segment, error)
}
