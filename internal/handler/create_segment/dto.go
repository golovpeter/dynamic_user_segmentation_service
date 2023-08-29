package create_segment

type CreateSegmentIn struct {
	SegmentSlug    string `json:"segment_slug" example:"AVITO_VOICE_MESSAGE" binding:"required"`
	PercentOfUsers int64  `json:"percent_users"`
}
