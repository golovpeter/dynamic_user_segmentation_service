package get_percent_segments

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/golovpeter/avito-trainee-task-2023/internal/cache/percent_segments"
	"github.com/golovpeter/avito-trainee-task-2023/internal/repository/segments"
)

type TestSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	mockSegmentsRepository *segments.MockRepository

	service GetPercentSegmentsService
}

func (ts *TestSuite) SetupTest() {
	ts.ctrl = gomock.NewController(ts.T())

	ts.mockSegmentsRepository = segments.NewMockRepository(ts.ctrl)

	ts.service = NewService(ts.mockSegmentsRepository)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

var (
	testPercentSegments = map[string]percent_segments.Segment{
		"AVITO_VOICE_MESSAGE": {
			Id:           1,
			Slug:         "AVITO_VOICE_MESSAGE",
			PercentUsers: 50,
		},
		"AVITO_DISCONT_30": {
			Id:           2,
			Slug:         "AVITO_DISCONT_30",
			PercentUsers: 30,
		},
	}
)

func (ts *TestSuite) Test_GetPercentSegments_Success() {
	ts.mockSegmentsRepository.EXPECT().
		GetPercentSegments().
		Times(1).
		Return(testPercentSegments, nil)

	percentSegments, err := ts.service.GetPercentSegments()

	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), testPercentSegments, percentSegments)
}

var testEmptyPercentSegments map[string]percent_segments.Segment

func (ts *TestSuite) Test_GetPercentSegments_RepositoryError() {
	ts.mockSegmentsRepository.EXPECT().
		GetPercentSegments().
		Times(1).
		Return(nil, errors.New("repository error"))

	percentSegments, err := ts.service.GetPercentSegments()

	assert.Error(ts.T(), err)
	assert.Equal(ts.T(), testEmptyPercentSegments, percentSegments)
}
