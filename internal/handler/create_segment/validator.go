package create_segment

import (
	"github.com/golovpeter/avito-trainee-task-2023/internal/common"
)

func validateInParams(in CreateSegmentIn) (bool, string, error) {
	if in.PercentOfUsers < 0 || in.PercentOfUsers > 100 {
		return false, "invalid percentage of users", nil
	}

	return common.ValidateSlug(in.SegmentSlug)
}
