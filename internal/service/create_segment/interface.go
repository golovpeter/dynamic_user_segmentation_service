package create_segment

type CreateSegmentService interface {
	CreateSegment(data *CreateSegmentData) error
}
