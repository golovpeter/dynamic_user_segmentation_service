package change_user_segments

import "time"

type ChangeUserSegmentsData struct {
	AddSegments    []string
	DeleteSegments []string
	UserID         int64
	ExpiredAt      time.Time
}
