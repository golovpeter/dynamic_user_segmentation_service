package get_user_segments

import (
	"math/rand"

	"github.com/golovpeter/avito-trainee-task-2023/internal/repository/user_segments"
)

type service struct {
	userSegmentsRepository user_segments.Repository
}

func NewService(
	userSegmentsRepository user_segments.Repository,
) *service {
	return &service{
		userSegmentsRepository: userSegmentsRepository,
	}
}

func (s *service) GetUserSegments(data *GetUserSegmentsData) ([]string, error) {
	allUserSegments, err := s.userSegmentsRepository.GetUserSegments(data.UserId)
	if err != nil {
		return []string{}, err
	}

	percentSegments := data.PercentSegmentsCache.Get()

	for slug, segmentInfo := range percentSegments {
		if _, ok := allUserSegments[slug]; ok {
			continue
		}

		var addSegment bool
		randomNumber := rand.Float64()

		if randomNumber <= float64(segmentInfo.PercentUsers)/100 {
			addSegment = true
		} else {
			addSegment = false
		}

		err = s.userSegmentsRepository.AddOneUserSegment(data.UserId, segmentInfo.Id, addSegment)
		if err != nil {
			return []string{}, nil
		}

		allUserSegments[slug] = user_segments.SegmentInfo{
			Slug:           slug,
			ID:             segmentInfo.Id,
			AddedToSegment: addSegment,
		}
	}

	var currentUserSegments = make([]string, 0, len(allUserSegments))

	for key, val := range allUserSegments {
		if val.AddedToSegment == true {
			currentUserSegments = append(currentUserSegments, key)
		}
	}

	return currentUserSegments, nil
}
