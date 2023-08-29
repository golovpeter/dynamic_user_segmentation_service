package create_segment

import "github.com/golovpeter/avito-trainee-task-2023/internal/repository/segments"

type service struct {
	repository segments.Repository
}

func NewService(
	segmentRepository segments.Repository,
) *service {
	return &service{
		repository: segmentRepository,
	}
}

func (s *service) CreateSegment(data *CreateSegmentData) error {
	return s.repository.CreateSegment(data.SegmentSlug, data.PercentOfUsers)
}
