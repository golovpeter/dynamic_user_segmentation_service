package delete_segment

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/golovpeter/avito-trainee-task-2023/internal/repository/segments"
)

type TestSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	mockSegmentsRepository *segments.MockRepository

	service DeleteSegmentService
}

func (ts *TestSuite) SetupTest() {
	ts.ctrl = gomock.NewController(ts.T())

	ts.mockSegmentsRepository = segments.NewMockRepository(ts.ctrl)

	ts.service = NewService(ts.mockSegmentsRepository)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

const testSlug = "AVITO_MESSENGER"

func (ts *TestSuite) Test_DeleteSegment_Success() {
	ts.mockSegmentsRepository.EXPECT().
		DeleteSegment(testSlug).
		Times(1).
		Return(int64(1), nil)

	err := ts.service.DeleteSegment(&DeleteSegmentData{
		SegmentSlug: testSlug,
	})

	assert.NoError(ts.T(), err)
}

func (ts *TestSuite) Test_DeleteSegment_ErrorSegmentNotFound() {
	ts.mockSegmentsRepository.EXPECT().
		DeleteSegment(testSlug).
		Times(1).
		Return(int64(0), nil)

	err := ts.service.DeleteSegment(&DeleteSegmentData{
		SegmentSlug: testSlug,
	})

	assert.ErrorIs(ts.T(), err, ErrSegmentNotFound)
}

func (ts *TestSuite) Test_DeleteSegment_OtherError() {
	ts.mockSegmentsRepository.EXPECT().
		DeleteSegment(testSlug).
		Times(1).
		Return(int64(0), errors.New("repository error"))

	err := ts.service.DeleteSegment(&DeleteSegmentData{
		SegmentSlug: testSlug,
	})

	assert.Error(ts.T(), err)
}
