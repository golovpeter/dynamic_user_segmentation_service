package get_user_segments

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

func (s *service) GetUserSegments(data *GetUserSegmentsData) ([]string, error) {
	userSegments, err := s.repository.GetUserSegments(data.UserId)
	if err != nil {
		return []string{}, err
	}

	return userSegments, nil
}
