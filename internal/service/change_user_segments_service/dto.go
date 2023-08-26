package change_user_segments_service

type ChangeUserSegmentsData struct {
	AddSegments    []string
	DeleteSegments []string
	UserID         int
}
