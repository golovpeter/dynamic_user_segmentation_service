package change_user_segments

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/golovpeter/avito-trainee-task-2023/internal/service/change_user_segments"
)

type TestSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	service *change_user_segments.MockChangeUserSegmentsService

	handler *handler

	ctx *gin.Context
}

func (ts *TestSuite) SetupTest() {
	ts.ctrl = gomock.NewController(ts.T())

	ts.service = change_user_segments.NewMockChangeUserSegmentsService(ts.ctrl)

	ts.handler = NewHandler(logrus.New(), ts.service)

	ts.ctx, _ = gin.CreateTestContext(httptest.NewRecorder())
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

const (
	testRequestSuccess = `
{
    "add_segments": [
        "AVITO_VOICE_MESSAGE",
		"AVITO_DISCOUNT_50"
    ],
	"delete_segments": [
		"AVITO_DISCOUNT_30",
		"AVITO_DISCOUNT_20"
    ],
    "user_id": 1003,
    "expired_at": "2023-09-28T23:41:00Z"
}
`
	testRequestInvalidJSON = `
{
	"invalid_json": "invalid json"
}
`
	testInvalidInParams = `
{
	"add_segments": [
		"AVITO_DISCOUNT_30",
		"AVITO_DISCOUNT_20"
    ],
	"delete_segments": [
		"AVITO_DISCOUNT_30",
		"AVITO_DISCOUNT_20"
    ],
    "user_id": 1003,
    "expired_at": "2023-09-28T23:41:00Z"
}
`
	testSegmentNotFound = `
{
	"add_segments": [
		"AVITO_DISCOUNT_30"
    ],
	"delete_segments": [],
    "user_id": 1003
}
`
)

var (
	testSegmentsData = &change_user_segments.ChangeUserSegmentsData{
		AddSegments: []string{
			"AVITO_VOICE_MESSAGE",
			"AVITO_DISCOUNT_50",
		},
		DeleteSegments: []string{
			"AVITO_DISCOUNT_30",
			"AVITO_DISCOUNT_20",
		},
		UserID:    1003,
		ExpiredAt: testTime,
	}

	testSegmentsDataNotFound = &change_user_segments.ChangeUserSegmentsData{
		AddSegments:    []string{"AVITO_DISCOUNT_30"},
		DeleteSegments: []string{},
		UserID:         1003,
	}

	testTime, _ = time.Parse(time.RFC3339, "2023-09-28T23:41:00Z")
)

func (ts *TestSuite) Test_ChangeUserSegment_Success() {
	body := io.NopCloser(bytes.NewReader([]byte(testRequestSuccess)))

	ts.ctx.Request = &http.Request{
		Method: http.MethodPost,
		Body:   body,
	}

	ts.service.EXPECT().
		ChangeUserSegments(testSegmentsData).
		Times(1).
		Return(nil)

	ts.handler.ChangeUserSegments(ts.ctx)
	assert.Equal(ts.T(), http.StatusOK, ts.ctx.Writer.Status())
}

func (ts *TestSuite) Test_ChangeUserSegment_ErrorBindingJSON() {
	body := io.NopCloser(bytes.NewReader([]byte(testRequestInvalidJSON)))

	ts.ctx.Request = &http.Request{
		Method: http.MethodPost,
		Body:   body,
	}

	ts.handler.ChangeUserSegments(ts.ctx)
	assert.Equal(ts.T(), http.StatusBadRequest, ts.ctx.Writer.Status())
}

func (ts *TestSuite) Test_ChangeUserSegment_invalidInParams() {
	body := io.NopCloser(bytes.NewReader([]byte(testInvalidInParams)))

	ts.ctx.Request = &http.Request{
		Method: http.MethodPost,
		Body:   body,
	}

	ts.handler.ChangeUserSegments(ts.ctx)
	assert.Equal(ts.T(), http.StatusBadRequest, ts.ctx.Writer.Status())
}

func (ts *TestSuite) Test_ChangeUserSegments_ErrorSegmentsNotFound() {
	body := io.NopCloser(bytes.NewReader([]byte(testSegmentNotFound)))

	ts.ctx.Request = &http.Request{
		Method: http.MethodPost,
		Body:   body,
	}

	ts.service.EXPECT().
		ChangeUserSegments(testSegmentsDataNotFound).
		Times(1).
		Return(change_user_segments.ErrorSegmentsNotFound{})

	ts.handler.ChangeUserSegments(ts.ctx)
	assert.Equal(ts.T(), http.StatusBadRequest, ts.ctx.Writer.Status())
}

func (ts *TestSuite) Test_ChangeUserSegments_OtherError() {
	body := io.NopCloser(bytes.NewReader([]byte(testRequestSuccess)))

	ts.ctx.Request = &http.Request{
		Method: http.MethodPost,
		Body:   body,
	}

	ts.service.EXPECT().
		ChangeUserSegments(testSegmentsData).
		Times(1).
		Return(errors.New("service error"))

	ts.handler.ChangeUserSegments(ts.ctx)
	assert.Equal(ts.T(), http.StatusInternalServerError, ts.ctx.Writer.Status())
}
