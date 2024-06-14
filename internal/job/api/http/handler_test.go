package http

import (
	"bytes"
	"encoding/json"
	"job-finder/internal/job/model"
	"job-finder/internal/job/usecase/mocks"
	"job-finder/pkg/response"
	"job-finder/pkg/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type JobHandlerTestSuite struct {
	suite.Suite
	mockUsecase *mocks.JobUsecase
	handler     *jobHandler
}

func (suite *JobHandlerTestSuite) SetupTest() {
	suite.mockUsecase = mocks.NewJobUsecase(suite.T())
	suite.handler = NewJobHandler(suite.mockUsecase)
}

func TestJobHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(JobHandlerTestSuite))
}

func (suite *JobHandlerTestSuite) prepareContext(method, path string, body any) (*gin.Context, *httptest.ResponseRecorder) {
	requestBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(method, path, bytes.NewBuffer(requestBody))
	c, _ := gin.CreateTestContext(w)
	c.Request = r

	return c, w
}

func (suite *JobHandlerTestSuite) TestPostNewJob() {
	req := &model.CreateJobReq{
		CompanyID:   uuid.New(),
		Title:       "Backend",
		Description: "Backend Developer with golang proficiency",
	}

	ctx, writer := suite.prepareContext(http.MethodPost, "/api/v1/jobs", req)

	suite.mockUsecase.On("Create", mock.Anything, req).
		Return(nil).Times(1)

	suite.handler.Create(ctx)

	suite.Equal(http.StatusOK, writer.Code)
}

func (suite *JobHandlerTestSuite) TestFindJobByKeyword() {
	query := map[string]interface{}{
		"keyword": "Backend",
	}
	ctx, writer := suite.prepareContext(http.MethodGet, "/api/v1/jobs", nil)
	ctx.Request.URL.RawQuery = "keyword=Backend"

	suite.mockUsecase.On("FindJob", mock.Anything, query).
		Return(
			[]*model.Job{{
				ID:          uuid.New(),
				CompanyID:   uuid.New(),
				Title:       "Backend",
				Description: "Backend Developer with golang proficency",
				CreatedAt:   time.Now(),
			}},
			nil,
		).Times(1)
	suite.handler.FindJob(ctx)

	res := &response.ResSuccess{}
	resBody := &model.Job{}
	json.Unmarshal(writer.Body.Bytes(), &res)
	utils.Copy(&res.Data, &resBody)

	suite.Equal(http.StatusOK, writer.Code)
	suite.Equal(query["keyword"], resBody.Title)
}
