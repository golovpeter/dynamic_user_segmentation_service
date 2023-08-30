package get_user_segments

//go:generate mockgen -destination=mocks.go -package=$GOPACKAGE -source=interfaces.go

type GetUserSegmentsService interface {
	GetUserSegments(data *GetUserSegmentsData) ([]string, error)
}
