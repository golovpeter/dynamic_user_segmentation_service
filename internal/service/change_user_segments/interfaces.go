package change_user_segments

//go:generate mockgen -destination=mocks.go -package=$GOPACKAGE -source=interfaces.go

type ChangeUserSegmentsService interface {
	ChangeUserSegments(data *ChangeUserSegmentsData) error
}
