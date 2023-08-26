package change_user_segments_service

import (
	"fmt"
)

type ErrorSegmentsNotFound struct {
	notFoundSlugs []string
}

func NewErrorSegmentsNotFound(notFoundSlugs []string) ErrorSegmentsNotFound {
	return ErrorSegmentsNotFound{
		notFoundSlugs: notFoundSlugs,
	}
}

func (e ErrorSegmentsNotFound) Error() string {
	return fmt.Sprintf("segments %v not found", e.notFoundSlugs)
}
