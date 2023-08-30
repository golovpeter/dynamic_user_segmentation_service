package create_segment

//go:generate mockgen -destination=mocks.go -package=$GOPACKAGE -source=interfaces.go

type CreateSegmentService interface {
	CreateSegment(data *CreateSegmentData) error
}
