package get_user_segments

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/golovpeter/avito-trainee-task-2023/internal/cache/percent_segments"
	"github.com/golovpeter/avito-trainee-task-2023/internal/repository/user_segments"
)

type TestSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	mockUserSegmentsRepository *user_segments.MockRepository

	service GetUserSegmentsService

	cache *percent_segments.Cache
}

func (ts *TestSuite) SetupTest() {
	ts.ctrl = gomock.NewController(ts.T())

	ts.mockUserSegmentsRepository = user_segments.NewMockRepository(ts.ctrl)

	ts.service = NewService(ts.mockUserSegmentsRepository)

	ts.cache = percent_segments.NewCache()
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

const testUserId int64 = 123

var (
	testUserSegments = []string{"AVITO_MESSENGER", "AVITO_PAY"}

	testMapSegments = map[string]user_segments.SegmentInfo{
		testUserSegments[0]: {
			Slug:           testUserSegments[0],
			ID:             1,
			AddedToSegment: true,
		},
		testUserSegments[1]: {
			Slug:           testUserSegments[1],
			ID:             2,
			AddedToSegment: true,
		},
	}
)

func (ts *TestSuite) Test_GetUserSegments_Success() {
	ts.mockUserSegmentsRepository.EXPECT().
		GetUserSegments(testUserId).
		Times(1).
		Return(testMapSegments, nil)

	userSegments, err := ts.service.GetUserSegments(&GetUserSegmentsData{
		UserId:               testUserId,
		PercentSegmentsCache: ts.cache,
	})

	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), testUserSegments, userSegments)
}

func (ts *TestSuite) Test_GetUserSegments_Error_Repository() {
	ts.mockUserSegmentsRepository.EXPECT().
		GetUserSegments(testUserId).
		Times(1).
		Return(nil, errors.New("repository error"))

	userSegments, err := ts.service.GetUserSegments(&GetUserSegmentsData{
		UserId: testUserId,
	})

	assert.Error(ts.T(), err)
	assert.Equal(ts.T(), []string{}, userSegments)
}
