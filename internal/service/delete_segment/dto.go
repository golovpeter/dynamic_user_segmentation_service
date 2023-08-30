package delete_segment

import "github.com/golovpeter/avito-trainee-task-2023/internal/cache/percent_segments"

type DeleteSegmentData struct {
	SegmentSlug          string
	PercentSegmentsCache *percent_segments.Cache
}
