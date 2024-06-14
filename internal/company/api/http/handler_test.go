package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"job-finder/internal/company/model"
	mocks "job-finder/internal/company/usecase/mocks"
	"job-finder/pkg/response"
	"job-finder/pkg/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CompanyHandlerTestSuite struct {
	suite.Suite
	mockService *mocks.CompanyUsecase
	handler     *companyHandler
}

func (suite *CompanyHandlerTestSuite) SetupTest() {
	suite.mockService = mocks.NewCompanyUsecase(suite.T())
	suite.handler = NewCompanyHandler(suite.mockService)
}

func TestCompanyHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CompanyHandlerTestSuite))
}

func (suite *CompanyHandlerTestSuite) prepareContext(method, path string, body any) (*gin.Context, *httptest.ResponseRecorder) {
	requestBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(method, path, bytes.NewBuffer(requestBody))
	c, _ := gin.CreateTestContext(w)
	c.Request = r

	return c, w
}

func (suite *CompanyHandlerTestSuite) TestRegisterSuccess() {
	req := &model.CreateCompany{
		Name: "Company 1",
	}

	ctx, writer := suite.prepareContext(http.MethodPost, "/api/v1/companies", req)

	suite.mockService.On("Register", mock.Anything, req).
		Return(
			&model.CompanyResponse{
				ID:   uuid.New(),
				Name: "Company 1",
			},
			nil,
		).Times(1)

	suite.handler.Create(ctx)

	res := &response.ResSuccess{}
	resBody := &model.CompanyResponse{}
	json.Unmarshal(writer.Body.Bytes(), &res)
	utils.Copy(&res.Data, &resBody)

	suite.Equal(http.StatusCreated, writer.Code)
	suite.Equal(req.Name, resBody.Name)
}

func (suite *CompanyHandlerTestSuite) TestRegisterFailed() {
	req := &model.CreateCompany{
		Name: "Company 1",
	}

	ctx, writer := suite.prepareContext(http.MethodPost, "/api/v1/companies", req)

	suite.mockService.On("Register", mock.Anything, req).
		Return(nil, errors.New("error")).Times(1)

	suite.handler.Create(ctx)

	res := &response.ResFailed{}
	json.Unmarshal(writer.Body.Bytes(), &res)
	suite.Equal(http.StatusInternalServerError, writer.Code)
	suite.Equal("error", res.Message)
}

func (suite *CompanyHandlerTestSuite) TestFindByID() {
	id := uuid.New()
	ctx, writer := suite.prepareContext(http.MethodGet, "/api/v1/companies/"+id.String(), nil)
	ctx.Params = gin.Params{gin.Param{Key: "id", Value: id.String()}}

	suite.mockService.On("FindByID", mock.Anything, id).
		Return(
			&model.Company{
				ID:   id,
				Name: "Company 1",
				Jobs: nil,
			},
			nil,
		).Times(1)

	suite.handler.FindByID(ctx)

	res := &response.ResSuccess{}
	resBody := &model.Company{}
	json.Unmarshal(writer.Body.Bytes(), &res)
	utils.Copy(&res.Data, &resBody)

	suite.Equal(http.StatusOK, writer.Code)
	suite.Equal(id, resBody.ID)
	suite.Equal("Company 1", resBody.Name)
}

func (suite *CompanyHandlerTestSuite) TestDeleteByID() {
	id := uuid.New()
	ctx, writer := suite.prepareContext(http.MethodDelete, "/api/v1/companies/"+id.String(), nil)
	ctx.Params = gin.Params{gin.Param{Key: "id", Value: id.String()}}

	suite.mockService.On("Delete", mock.Anything, id).
		Return(nil).Times(1)

	suite.handler.DeleteByID(ctx)

	suite.Equal(http.StatusOK, writer.Code)
}
