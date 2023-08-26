package delete_segment

import "github.com/golovpeter/avito-trainee-task-2023/internal/repository/segments"

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

func (s *service) DeleteSegment(data *DeleteSegmentData) error {
	deleted, err := s.repository.DeleteSegment(data.SegmentSlug)
	if err != nil {
		return err
	}

	if !deleted {
		return ErrSegmentNotFound
	}

	return err
}
