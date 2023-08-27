package user_segments

import "time"

type ChangeUserSegmentsData struct {
	AddSegmentsIds    []int64
	DeleteSegmentsIds []int64
	UserID            int64
	ExpiredAt         time.Time
}
