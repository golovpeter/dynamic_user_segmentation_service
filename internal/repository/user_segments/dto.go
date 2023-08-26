package user_segments

type ChangeUserSegmentsData struct {
	AddSegmentsIds    []int64
	DeleteSegmentsIds []int64
	UserID            int64
}
