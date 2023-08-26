package create_segment

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

	service CreateSegmentService
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

func (ts *TestSuite) Test_CreateSegment_Success() {
	ts.mockSegmentsRepository.EXPECT().
		CreateSegment(testSlug).
		Times(1).
		Return(nil)

	err := ts.service.CreateSegment(&CreateSegmentData{
		SegmentSlug: testSlug,
	})

	assert.NoError(ts.T(), err)
}

func (ts *TestSuite) Test_CreateSegment_RepositoryError() {
	ts.mockSegmentsRepository.EXPECT().
		CreateSegment(testSlug).
		Times(1).
		Return(errors.New("repository error"))

	err := ts.service.CreateSegment(&CreateSegmentData{
		SegmentSlug: testSlug,
	})

	assert.Error(ts.T(), err)
}
