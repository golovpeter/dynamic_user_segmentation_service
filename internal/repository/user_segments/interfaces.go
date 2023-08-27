package user_segments

//go:generate mockgen -destination=mocks.go -package=$GOPACKAGE -source=interfaces.go

type Repository interface {
	ChangeUserSegments(changeData ChangeUserSegmentsData) error
	DeleteUsersAfterTime() error
}
