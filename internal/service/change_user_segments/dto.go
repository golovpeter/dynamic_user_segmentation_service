package change_user_segments

type ChangeUserSegmentsData struct {
	AddSegments    []string
	DeleteSegments []string
	UserID         int64
}
