package create_segment

type CreateSegmentIn struct {
	SegmentSlug string `json:"segment_slug" example:"AVITO_VOICE_MESSAGE" binding:"required"`
	// required: false
	// description: PercentOfUsers is an optional parameter.
	PercentOfUsers int64 `json:"percent_users" example:"50"`
}
