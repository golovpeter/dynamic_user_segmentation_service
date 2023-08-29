package get_user_segments

import "github.com/golovpeter/avito-trainee-task-2023/internal/repository/segments"

type GetUserSegmentsData struct {
	UserId          int64
	PercentSegments map[string]segments.Segment
}
