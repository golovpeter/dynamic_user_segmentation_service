package change_user_segments_service

type ChangeUserSegmentsService interface {
	ChangeUserSegments(data *ChangeUserSegmentsData) error
}
