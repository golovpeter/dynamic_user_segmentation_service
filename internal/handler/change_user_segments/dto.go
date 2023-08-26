package change_user_segments

type ChangeUserSegmentsIn struct {
	AddSegments    []string `json:"add_segments"`
	DeleteSegments []string `json:"delete_segments"`
	UserID         int      `json:"user_id"`
}
