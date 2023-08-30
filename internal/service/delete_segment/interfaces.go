package delete_segment

//go:generate mockgen -destination=mocks.go -package=$GOPACKAGE -source=interfaces.go

type DeleteSegmentService interface {
	DeleteSegment(data *DeleteSegmentData) error
}
