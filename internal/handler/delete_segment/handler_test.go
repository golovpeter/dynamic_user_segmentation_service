package delete_segment

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

	"github.com/golovpeter/avito-trainee-task-2023/internal/cache/percent_segments"
	delete_segment_service "github.com/golovpeter/avito-trainee-task-2023/internal/service/delete_segment"
)

type TestSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	service *delete_segment_service.MockDeleteSegmentService

	handler *handler

	cache *percent_segments.Cache

	ctx *gin.Context
}

func (ts *TestSuite) SetupTest() {
	ts.ctrl = gomock.NewController(ts.T())

	ts.service = delete_segment_service.NewMockDeleteSegmentService(ts.ctrl)

	ts.cache = percent_segments.NewCache()

	ts.handler = NewHandler(logrus.New(), ts.service, ts.cache)

	ts.ctx, _ = gin.CreateTestContext(httptest.NewRecorder())
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

const (
	testSuccessRequest = `
{
	"segment_slug": "AVITO_VOICE_MESSAGE"
}
`
	testInvalidJSONRequest = `
{
	"invalid_json": "invalid_json""
}
`
	testInvalidInParamsRequest = `
{
	"segment_slug": "avito_voice_message"
}
`
)

func (ts *TestSuite) Test_DeleteSegment_Success() {
	body := io.NopCloser(bytes.NewReader([]byte(testSuccessRequest)))

	ts.ctx.Request = &http.Request{
		Method: http.MethodPost,
		Body:   body,
	}

	ts.service.EXPECT().
		DeleteSegment(&delete_segment_service.DeleteSegmentData{
			SegmentSlug:          "AVITO_VOICE_MESSAGE",
			PercentSegmentsCache: ts.cache,
		}).
		Times(1).
		Return(nil)

	ts.handler.DeleteSegment(ts.ctx)
	assert.Equal(ts.T(), http.StatusOK, ts.ctx.Writer.Status())

}

func (ts *TestSuite) Test_DeleteSegment_InvalidJSON() {
	body := io.NopCloser(bytes.NewReader([]byte(testInvalidJSONRequest)))

	ts.ctx.Request = &http.Request{
		Method: http.MethodPost,
		Body:   body,
	}

	ts.handler.DeleteSegment(ts.ctx)
	assert.Equal(ts.T(), http.StatusBadRequest, ts.ctx.Writer.Status())
}

func (ts *TestSuite) Test_DeleteSegment_InvalidInParams() {
	body := io.NopCloser(bytes.NewReader([]byte(testInvalidInParamsRequest)))

	ts.ctx.Request = &http.Request{
		Method: http.MethodPost,
		Body:   body,
	}

	ts.handler.DeleteSegment(ts.ctx)
	assert.Equal(ts.T(), http.StatusBadRequest, ts.ctx.Writer.Status())
}

func (ts *TestSuite) Test_DeleteSegment_ErrorSegmentNotFound() {
	body := io.NopCloser(bytes.NewReader([]byte(testSuccessRequest)))

	ts.ctx.Request = &http.Request{
		Method: http.MethodPost,
		Body:   body,
	}

	ts.service.EXPECT().DeleteSegment(&delete_segment_service.DeleteSegmentData{
		SegmentSlug:          "AVITO_VOICE_MESSAGE",
		PercentSegmentsCache: ts.cache,
	}).Times(1).
		Return(delete_segment_service.ErrSegmentNotFound)

	ts.handler.DeleteSegment(ts.ctx)
	assert.Equal(ts.T(), http.StatusBadRequest, ts.ctx.Writer.Status())
}

func (ts *TestSuite) Test_DeleteSegment_OtherError() {
	body := io.NopCloser(bytes.NewReader([]byte(testSuccessRequest)))

	ts.ctx.Request = &http.Request{
		Method: http.MethodPost,
		Body:   body,
	}

	ts.service.EXPECT().DeleteSegment(&delete_segment_service.DeleteSegmentData{
		SegmentSlug:          "AVITO_VOICE_MESSAGE",
		PercentSegmentsCache: ts.cache,
	}).Times(1).
		Return(errors.New("service error"))

	ts.handler.DeleteSegment(ts.ctx)
	assert.Equal(ts.T(), http.StatusInternalServerError, ts.ctx.Writer.Status())
}
