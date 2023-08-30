package get_percent_segments

import (
	"github.com/golovpeter/avito-trainee-task-2023/internal/cache/percent_segments"
	"github.com/golovpeter/avito-trainee-task-2023/internal/repository/segments"
)

type service struct {
	repository segments.Repository
}

func NewService(
	repository segments.Repository,
) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetPercentSegments() (map[string]percent_segments.Segment, error) {
	return s.repository.GetPercentSegments()
}
