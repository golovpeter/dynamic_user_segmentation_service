package user_segments

import "time"

type ChangeUserSegmentsData struct {
	AddSegmentsIds    []int64
	DeleteSegmentsIds []int64
	UserID            int64
	ExpiredAt         time.Time
}

type SegmentInfo struct {
	Slug           string `db:"slug"`
	ID             int    `db:"segment_id"`
	AddedToSegment bool   `db:"added_to_segment"`
}
