package create_segment

import (
	"github.com/golovpeter/avito-trainee-task-2023/internal/common"
)

func validateInParams(slug string) (bool, string, error) {
	return common.ValidateSlug(slug)
}
