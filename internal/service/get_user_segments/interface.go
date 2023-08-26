package get_user_segments

type GetUserSegmentsService interface {
	GetUserSegments(data *GetUserSegmentsData) ([]string, error)
}
