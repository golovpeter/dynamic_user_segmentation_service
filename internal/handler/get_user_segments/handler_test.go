package get_user_segments

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/golovpeter/avito-trainee-task-2023/internal/cache/percent_segments"
	"github.com/golovpeter/avito-trainee-task-2023/internal/service/get_user_segments"
)

type TestSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	service *get_user_segments.MockGetUserSegmentsService

	handler *handler

	cache *percent_segments.Cache

	ctx *gin.Context
}

func (ts *TestSuite) SetupTest() {
	ts.ctrl = gomock.NewController(ts.T())

	ts.service = get_user_segments.NewMockGetUserSegmentsService(ts.ctrl)

	ts.cache = percent_segments.NewCache()

	ts.handler = NewHandler(logrus.New(), ts.service, ts.cache)

	ts.ctx, _ = gin.CreateTestContext(httptest.NewRecorder())
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

var (
	testUserId        int64  = 1000
	testInvalidUserId string = "invalidUserId"
)

func (ts *TestSuite) Test_GetUserSegment_Success() {
	ts.ctx.Request = &http.Request{
		Method: http.MethodGet,
	}

	ts.ctx.Params = []gin.Param{
		{
			Key:   "user_id",
			Value: "1000",
		},
	}

	ts.service.EXPECT().
		GetUserSegments(&get_user_segments.GetUserSegmentsData{
			UserId:               testUserId,
			PercentSegmentsCache: ts.cache,
		}).
		Times(1).
		Return([]string{}, nil)

	ts.handler.GetUserSegments(ts.ctx)

	assert.Equal(ts.T(), http.StatusOK, ts.ctx.Writer.Status())
}

func (ts *TestSuite) Test_GetUserSegments_InvalidID() {
	ts.ctx.Request = &http.Request{
		Method: http.MethodGet,
	}

	ts.ctx.Params = []gin.Param{
		{
			Key:   "user_id",
			Value: testInvalidUserId,
		},
	}

	ts.handler.GetUserSegments(ts.ctx)
	assert.Equal(ts.T(), http.StatusBadRequest, ts.ctx.Writer.Status())
}

func (ts *TestSuite) Test_GetUserSegments_ServiceError() {
	ts.ctx.Request = &http.Request{
		Method: http.MethodGet,
	}

	ts.ctx.Params = []gin.Param{
		{
			Key:   "user_id",
			Value: "1000",
		},
	}

	ts.service.EXPECT().
		GetUserSegments(&get_user_segments.GetUserSegmentsData{
			UserId:               testUserId,
			PercentSegmentsCache: ts.cache,
		}).
		Times(1).
		Return([]string{}, errors.New("service error"))

	ts.handler.GetUserSegments(ts.ctx)
	assert.Equal(ts.T(), http.StatusInternalServerError, ts.ctx.Writer.Status())
}
