package change_user_segments

import "time"

type ChangeUserSegmentsIn struct {
	// required: false
	// description: AddSegments is an optional parameter.
	AddSegments []string `json:"add_segments" example:"AVITO_VOICE_MESSAGE,AVITO_DICSOUNT_30"`
	// required: false
	// description: DeleteSegments is an optional parameter.
	DeleteSegments []string `json:"delete_segments" example:"AVITO_DICOUNT_50"`
	UserID         int64    `json:"user_id" example:"1000" binding:"required"`
	// required: false
	// description: ExpiredAt is an optional parameter.
	ExpiredAt time.Time `json:"expired_at,omitempty" example:"2023-08-27T15:40:00Z"`
}
