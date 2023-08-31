package create_segment

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	create_segment_service "github.com/golovpeter/avito-trainee-task-2023/internal/service/create_segment"
)

type TestSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	service *create_segment_service.MockCreateSegmentService

	handler *handler

	ctx *gin.Context
}

func (ts *TestSuite) SetupTest() {
	ts.ctrl = gomock.NewController(ts.T())

	ts.service = create_segment_service.NewMockCreateSegmentService(ts.ctrl)

	ts.handler = NewHandler(logrus.New(), ts.service)

	ts.ctx, _ = gin.CreateTestContext(httptest.NewRecorder())
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

var (
	testSuccessRequest = `
{
	"segment_slug": "AVITO_VOICE_MESSAGE",
	"percent_users": 50
}
`
	testInvalidJSON = `
{
	"invalid_json": "invalid_json"
}`

	testInvalidInParams = `
{
	"segment_slug": "avito_voice_message",
	"percent_users": -100
}
`
)

var (
	testSuccessData = &create_segment_service.CreateSegmentData{
		SegmentSlug:    "AVITO_VOICE_MESSAGE",
		PercentOfUsers: 50,
	}
)

func (ts *TestSuite) Test_CreateSegment_Success() {
	body := io.NopCloser(bytes.NewReader([]byte(testSuccessRequest)))

	ts.ctx.Request = &http.Request{
		Method: http.MethodPost,
		Body:   body,
	}

	ts.service.EXPECT().
		CreateSegment(testSuccessData).
		Times(1).
		Return(nil)

	ts.handler.CreateSegment(ts.ctx)
	assert.Equal(ts.T(), http.StatusOK, ts.ctx.Writer.Status())
}

func (ts *TestSuite) Test_CreateSegment_InvalidJSON() {
	body := io.NopCloser(bytes.NewReader([]byte(testInvalidJSON)))

	ts.ctx.Request = &http.Request{
		Method: http.MethodPost,
		Body:   body,
	}

	ts.handler.CreateSegment(ts.ctx)
	assert.Equal(ts.T(), http.StatusBadRequest, ts.ctx.Writer.Status())
}

func (ts *TestSuite) Test_CreateSegment_InvalidInParams() {
	body := io.NopCloser(bytes.NewReader([]byte(testInvalidInParams)))

	ts.ctx.Request = &http.Request{
		Method: http.MethodPost,
		Body:   body,
	}

	ts.handler.CreateSegment(ts.ctx)
	assert.Equal(ts.T(), http.StatusBadRequest, ts.ctx.Writer.Status())
}

func (ts *TestSuite) Test_CreateSegment_ServiceError() {
	body := io.NopCloser(bytes.NewReader([]byte(testSuccessRequest)))

	ts.ctx.Request = &http.Request{
		Method: http.MethodPost,
		Body:   body,
	}

	ts.service.EXPECT().
		CreateSegment(testSuccessData).
		Times(1).
		Return(errors.New("service error"))

	ts.handler.CreateSegment(ts.ctx)
	assert.Equal(ts.T(), http.StatusInternalServerError, ts.ctx.Writer.Status())
}
