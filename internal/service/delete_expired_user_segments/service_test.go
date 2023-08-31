package delete_expired_user_segments

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/golovpeter/avito-trainee-task-2023/internal/repository/user_segments"
)

type TestSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	mockSegmentsRepository *user_segments.MockRepository

	service DeleteExpiredUserSegmentsService
}

func (ts *TestSuite) SetupTest() {
	ts.ctrl = gomock.NewController(ts.T())

	ts.mockSegmentsRepository = user_segments.NewMockRepository(ts.ctrl)

	ts.service = NewService(ts.mockSegmentsRepository)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (ts *TestSuite) Test_DeleteExpiredUserSegments_Success() {
	ts.mockSegmentsRepository.EXPECT().
		DeleteExpiredUserSegments().
		Times(1).
		Return(nil)

	err := ts.service.DeleteExpiredUserSegments()

	assert.NoError(ts.T(), err)
}

func (ts *TestSuite) Test_DeleteExpiredUserSegments_RepositoryError() {
	ts.mockSegmentsRepository.EXPECT().
		DeleteExpiredUserSegments().
		Times(1).
		Return(errors.New("repository error"))

	err := ts.service.DeleteExpiredUserSegments()

	assert.Error(ts.T(), err)
}
