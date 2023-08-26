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

func (s *service) DeleteSegmentService(data *DeleteSegmentData) error {
	rowsAffected, err := s.repository.DeleteSegment(data.SegmentSlug)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrSegmentNotFound
	}

	return err
}
