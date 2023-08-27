package change_user_segments

type ChangeUserSegmentsIn struct {
	AddSegments    []string `json:"add_segments" example:"AVITO_VOICE_MESSAGE,AVITO_DICSOUNT_30"`
	DeleteSegments []string `json:"delete_segments" example:"AVITO_DICOUNT_50"`
	UserID         int64    `json:"user_id" example:"1000"`
}
