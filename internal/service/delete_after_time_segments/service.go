package delete_after_time_segments

import "github.com/golovpeter/avito-trainee-task-2023/internal/repository/user_segments"

type service struct {
	repository user_segments.Repository
}

func NewService(
	repository user_segments.Repository,
) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) DeleteUsersAfterTime() error {
	return s.repository.DeleteUsersAfterTime()
}
