package change_user_segments

import (
	"errors"
	"github.com/golovpeter/avito-trainee-task-2023/internal/repository/user_segments"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/golovpeter/avito-trainee-task-2023/internal/repository/segments"
)

type TestSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	mockSegmentsRepository        *segments.MockRepository
	mockUsersToSegmentsRepository *user_segments.MockRepository

	service ChangeUserSegmentsService
}

func (ts *TestSuite) SetupTest() {
	ts.ctrl = gomock.NewController(ts.T())

	ts.mockSegmentsRepository = segments.NewMockRepository(ts.ctrl)
	ts.mockUsersToSegmentsRepository = user_segments.NewMockRepository(ts.ctrl)

	ts.service = NewService(
		ts.mockSegmentsRepository,
		ts.mockUsersToSegmentsRepository,
	)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

var (
	testSlugsToAdd    = []string{"AVITO_MESSENGER", "AVITO_PAY"}
	testSlugsIdsToAdd = []int64{1, 2}
	testSlugsToAddMap = map[string]int64{
		testSlugsToAdd[0]: testSlugsIdsToAdd[0],
		testSlugsToAdd[1]: testSlugsIdsToAdd[1],
	}

	testSlugsToDelete    = []string{"AVITO_MAIN_PAGE", "AVITO_BANK"}
	testSlugsIdsToDelete = []int64{3, 4}
	testSlugsToDeleteMap = map[string]int64{
		testSlugsToDelete[0]: testSlugsIdsToDelete[0],
		testSlugsToDelete[1]: testSlugsIdsToDelete[1],
	}

	testUserId int64 = 123
)

func (ts *TestSuite) Test_ChangeUserSegments_AddSegments_Success() {
	ts.mockSegmentsRepository.EXPECT().
		GetActiveSegmentsIdsBySlugs(gomock.InAnyOrder(testSlugsToAdd)).
		Times(1).
		Return(testSlugsToAddMap, nil)

	ts.mockUsersToSegmentsRepository.EXPECT().
		ChangeUserSegments(user_segments.ChangeUserSegmentsData{
			AddSegmentsIds:    testSlugsIdsToAdd,
			DeleteSegmentsIds: make([]int64, 0),
			UserID:            testUserId,
		}).
		Times(1).
		Return(nil)

	err := ts.service.ChangeUserSegments(&ChangeUserSegmentsData{
		AddSegments: testSlugsToAdd,
		UserID:      testUserId,
	})

	assert.NoError(ts.T(), err)
}

func (ts *TestSuite) Test_ChangeUserSegments_DeleteSegments_Success() {
	ts.mockSegmentsRepository.EXPECT().
		GetActiveSegmentsIdsBySlugs(gomock.InAnyOrder(testSlugsToDelete)).
		Times(1).
		Return(testSlugsToDeleteMap, nil)

	ts.mockUsersToSegmentsRepository.EXPECT().
		ChangeUserSegments(user_segments.ChangeUserSegmentsData{
			AddSegmentsIds:    make([]int64, 0),
			DeleteSegmentsIds: testSlugsIdsToDelete,
			UserID:            testUserId,
		}).
		Times(1).
		Return(nil)

	err := ts.service.ChangeUserSegments(&ChangeUserSegmentsData{
		DeleteSegments: testSlugsToDelete,
		UserID:         testUserId,
	})

	assert.NoError(ts.T(), err)
}

func (ts *TestSuite) Test_ChangeUserSegments_AddAndDeleteSegments_Success() {
	mergedSlugs := append(testSlugsToAdd, testSlugsToDelete...)
	mergedSlugIds := append(testSlugsIdsToAdd, testSlugsIdsToDelete...)

	mergedSlugsIdsMap := make(map[string]int64, len(mergedSlugs))
	for i := range mergedSlugs {
		mergedSlugsIdsMap[mergedSlugs[i]] = mergedSlugIds[i]
	}

	ts.mockSegmentsRepository.EXPECT().
		GetActiveSegmentsIdsBySlugs(gomock.InAnyOrder(mergedSlugs)).
		Times(1).
		Return(mergedSlugsIdsMap, nil)

	ts.mockUsersToSegmentsRepository.EXPECT().
		ChangeUserSegments(user_segments.ChangeUserSegmentsData{
			AddSegmentsIds:    testSlugsIdsToAdd,
			DeleteSegmentsIds: testSlugsIdsToDelete,
			UserID:            testUserId,
		}).
		Times(1).
		Return(nil)

	err := ts.service.ChangeUserSegments(&ChangeUserSegmentsData{
		AddSegments:    testSlugsToAdd,
		DeleteSegments: testSlugsToDelete,
		UserID:         testUserId,
	})

	assert.NoError(ts.T(), err)
}

func (ts *TestSuite) Test_ChangeUserSegments_AddSegments_Error_UserSegmentsRepository() {
	ts.mockSegmentsRepository.EXPECT().
		GetActiveSegmentsIdsBySlugs(gomock.InAnyOrder(testSlugsToAdd)).
		Times(1).
		Return(testSlugsToAddMap, nil)

	ts.mockUsersToSegmentsRepository.EXPECT().
		ChangeUserSegments(user_segments.ChangeUserSegmentsData{
			AddSegmentsIds:    testSlugsIdsToAdd,
			DeleteSegmentsIds: make([]int64, 0),
			UserID:            testUserId,
		}).
		Times(1).
		Return(errors.New("repository error"))

	err := ts.service.ChangeUserSegments(&ChangeUserSegmentsData{
		AddSegments: testSlugsToAdd,
		UserID:      testUserId,
	})

	assert.Error(ts.T(), err)
}

func (ts *TestSuite) Test_ChangeUserSegments_AddSegments_Error_SlugNotFound() {
	notFullUserSegmentsMap := map[string]int64{
		testSlugsToAdd[0]: testSlugsIdsToAdd[0],
	}

	ts.mockSegmentsRepository.EXPECT().
		GetActiveSegmentsIdsBySlugs(gomock.InAnyOrder(testSlugsToAdd)).
		Times(1).
		Return(notFullUserSegmentsMap, nil)

	err := ts.service.ChangeUserSegments(&ChangeUserSegmentsData{
		AddSegments: testSlugsToAdd,
		UserID:      testUserId,
	})

	assert.EqualError(ts.T(), err, NewErrorSegmentsNotFound([]string{testSlugsToAdd[1]}).Error())
}

func (ts *TestSuite) Test_ChangeUserSegments_AddSegments_Error_SlugRepository() {
	ts.mockSegmentsRepository.EXPECT().
		GetActiveSegmentsIdsBySlugs(gomock.InAnyOrder(testSlugsToAdd)).
		Times(1).
		Return(nil, errors.New("repository error"))

	err := ts.service.ChangeUserSegments(&ChangeUserSegmentsData{
		AddSegments: testSlugsToAdd,
		UserID:      testUserId,
	})

	assert.Error(ts.T(), err)
}
