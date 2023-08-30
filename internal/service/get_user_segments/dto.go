package get_user_segments

import (
	"github.com/golovpeter/avito-trainee-task-2023/internal/cache/percent_segments"
)

type GetUserSegmentsData struct {
	UserId               int64
	PercentSegmentsCache *percent_segments.Cache
}
