package change_user_segments_service

import (
	"github.com/golovpeter/avito-trainee-task-2023/internal/repository/segments"
	"github.com/golovpeter/avito-trainee-task-2023/internal/repository/user_segments"
)

type service struct {
	segmentsRepository     segments.Repository
	userSegmentsRepository user_segments.Repository
}

func NewService(
	segmentsRepository segments.Repository,
	userSegmentsRepository user_segments.Repository,

) *service {
	return &service{
		segmentsRepository:     segmentsRepository,
		userSegmentsRepository: userSegmentsRepository,
	}
}

func (s *service) ChangeUserSegments(data *ChangeUserSegmentsData) error {
	uniqueSlugs := s.getUniqueSlugs(data)

	slugsIds, err := s.segmentsRepository.GetActiveSegmentsIdsBySlugs(uniqueSlugs)
	if err != nil {
		return err
	}

	if len(uniqueSlugs) != len(slugsIds) {
		var notFoundSlugs []string

		for _, slug := range uniqueSlugs {
			if _, ok := slugsIds[slug]; !ok {
				notFoundSlugs = append(notFoundSlugs, slug)
			}
		}

		return NewErrorSegmentsNotFound(notFoundSlugs)
	}

	addSegmentsIds := make([]int64, 0, len(data.AddSegments))
	deleteSegmentsIds := make([]int64, 0, len(data.DeleteSegments))

	for _, segment := range data.AddSegments {
		if _, ok := slugsIds[segment]; ok {
			addSegmentsIds = append(addSegmentsIds, slugsIds[segment])
		}
	}

	for _, segment := range data.DeleteSegments {
		if _, ok := slugsIds[segment]; ok {
			deleteSegmentsIds = append(deleteSegmentsIds, slugsIds[segment])
		}
	}

	err = s.userSegmentsRepository.ChangeUserSegments(
		user_segments.ChangeUserSegmentsData{
			AddSegmentsIds:    addSegmentsIds,
			DeleteSegmentsIds: deleteSegmentsIds,
			UserID:            data.UserID,
		},
	)

	return err
}

func (s *service) getUniqueSlugs(data *ChangeUserSegmentsData) []string {
	slugs := make(map[string]struct{})

	for _, slug := range data.AddSegments {
		slugs[slug] = struct{}{}
	}

	for _, slug := range data.DeleteSegments {
		slugs[slug] = struct{}{}
	}

	var uniqueSlugs []string
	for key, _ := range slugs {
		uniqueSlugs = append(uniqueSlugs, key)
	}

	return uniqueSlugs
}
