package user_segments

type Repository interface {
	ChangeUserSegments(changeData ChangeUserSegmentsData) error
}
