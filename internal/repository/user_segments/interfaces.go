package user_segments

//go:generate mockgen -destination=mocks.go -package=$GOPACKAGE -source=interfaces.go

type Repository interface {
	ChangeUserSegments(changeData ChangeUserSegmentsData) error
	DeleteExpiredUserSegments() error
	GetUserSegments(id int64) (map[string]SegmentInfo, error)
	AddOneUserSegment(userId, segmentId int64, addedSegment bool) error
}
