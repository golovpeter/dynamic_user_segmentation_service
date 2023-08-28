package delete_segment

type DeleteSegmentIn struct {
	SegmentSlug string `json:"segment_slug" example:"AVITO_VOICE_MESSAGE" binding:"required"`
}
