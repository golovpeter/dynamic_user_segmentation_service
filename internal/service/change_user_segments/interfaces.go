package change_user_segments

type ChangeUserSegmentsService interface {
	ChangeUserSegments(data *ChangeUserSegmentsData) error
}
