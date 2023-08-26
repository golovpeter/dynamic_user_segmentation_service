package delete_segment

type DeleteSegmentService interface {
	DeleteSegment(data *DeleteSegmentData) error
}
