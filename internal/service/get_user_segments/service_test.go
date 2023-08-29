package get_user_segments

import (
	"errors"
	"github.com/golovpeter/avito-trainee-task-2023/internal/repository/user_segments"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type TestSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	mockSegmentsRepository *user_segments.MockRepository

	service GetUserSegmentsService
}

func (ts *TestSuite) SetupTest() {
	ts.ctrl = gomock.NewController(ts.T())

	ts.mockSegmentsRepository = user_segments.NewMockRepository(ts.ctrl)

	ts.service = NewService(ts.mockSegmentsRepository)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

const testUserId int64 = 123

var testUserSegments = []string{"AVITO_MESSENGER, AVITO_PAY"}

func (ts *TestSuite) Test_GetUserSegments_Success() {
	ts.mockSegmentsRepository.EXPECT().
		GetUserSegments(testUserId).
		Times(1).
		Return(testUserSegments, nil)

	userSegments, err := ts.service.GetUserSegments(&GetUserSegmentsData{
		UserId: testUserId,
	})

	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), testUserSegments, userSegments)
}

func (ts *TestSuite) Test_GetUserSegments_Error_Repository() {
	ts.mockSegmentsRepository.EXPECT().
		GetUserSegments(testUserId).
		Times(1).
		Return(nil, errors.New("repository error"))

	userSegments, err := ts.service.GetUserSegments(&GetUserSegmentsData{
		UserId: testUserId,
	})

	assert.Error(ts.T(), err)
	assert.Equal(ts.T(), []string{}, userSegments)
}
