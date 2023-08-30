package delete_expired_user_segments

//go:generate mockgen -destination=mocks.go -package=$GOPACKAGE -source=interfaces.go

type DeleteExpiredUserSegmentsService interface {
	DeleteExpiredUserSegments() error
}
