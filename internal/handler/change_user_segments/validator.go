package change_user_segments

import (
	"fmt"
)

func validateInParams(in *ChangeUserSegmentsIn) (bool, string) {
	if in.UserID <= 0 {
		return false, "userID should be positive number"
	}

	if len(in.AddSegments) == 0 && len(in.DeleteSegments) == 0 {
		return false, "at least one add or delete segments shouldn't be empty"
	}

	interSegments := intersectionSegments(in.AddSegments, in.DeleteSegments)

	if len(interSegments) > 0 {
		return false, fmt.Sprintf("segments %v present in both arrays", interSegments)
	}

	return true, ""
}

func intersectionSegments(addSegments, deleteSegments []string) []string {
	var intersValues []string
	checkIntersValues := make(map[string]bool)

	for _, addSegment := range addSegments {
		checkIntersValues[addSegment] = true
	}

	for _, deleteSegment := range deleteSegments {
		if _, ok := checkIntersValues[deleteSegment]; ok {
			intersValues = append(intersValues, deleteSegment)
		}
	}

	return intersValues
}
