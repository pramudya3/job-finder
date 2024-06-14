package http

import (
	"errors"
	"job-finder/internal/company/model"
	"job-finder/internal/company/usecase"
	"job-finder/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type companyHandler struct {
	companyUsecase usecase.CompanyUsecase
}

func NewCompanyHandler(companyUc usecase.CompanyUsecase) *companyHandler {
	return &companyHandler{
		companyUsecase: companyUc,
	}
}

// Register godoc
//
//	@Summary	Register company
//	@Tags		companies
//	@Produce	json
//	@Param		_	body		model.CreateCompany	true	"Body"
//	@Success	201	{object}	response.ResSuccess
//	@Failure	400	{object}	response.ResFailed
//	@Failure	404	{object}	response.ResFailed
//	@Failure	500	{object}	response.ResFailed
//	@Router		/companies [post]
func (h *companyHandler) Create(c *gin.Context) {
	payload := &model.CreateCompany{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseFailed(err.Error()))
		return
	}

	res, err := h.companyUsecase.Register(c, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseFailed(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.ResponseSuccess(res, nil))
}

// DeleteCompanyByID godoc
//
//	@Summary	Delete company by id
//	@Tags		companies
//	@Produce	json
//	@Param		id	path	string	true	"Company ID"
//	@Success	200
//	@Failure	400	{object}	response.ResFailed
//	@Failure	404	{object}	response.ResFailed
//	@Failure	500	{object}	response.ResFailed
//	@Router		/companies/{id} [delete]
func (h *companyHandler) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	companyID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseFailed(err.Error()))
		return
	}

	if err := h.companyUsecase.Delete(c, companyID); err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Company not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, response.ResponseFailed(err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

// FindCompanyByID godoc
//
//	@Summary	Find company by id
//	@Tags		companies
//	@Produce	json
//	@Param		id	path		string	true	"Company ID"
//	@Success	200	{object}	response.ResSuccess
//	@Failure	400	{object}	response.ResFailed
//	@Failure	404	{object}	response.ResFailed
//	@Failure	500	{object}	response.ResFailed
//	@Router		/companies/{id} [get]
func (h *companyHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	companyID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseFailed(err.Error()))
		return
	}

	company, err := h.companyUsecase.FindByID(c, companyID)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Company not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, response.ResponseFailed(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.ResponseSuccess(company, nil))
}
