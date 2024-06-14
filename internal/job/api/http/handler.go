package http

import (
	"job-finder/internal/job/model"
	"job-finder/internal/job/usecase"
	response "job-finder/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type jobHandler struct {
	jobUsecase usecase.JobUsecase
}

func NewJobHandler(jobUsecase usecase.JobUsecase) *jobHandler {
	return &jobHandler{
		jobUsecase: jobUsecase,
	}
}

// CreateNewJob godoc
//
//	@Summary	Create new job
//	@Tags		jobs
//	@Produce	json
//	@Param		_	body	model.CreateJobReq	true	"Body"
//	@Success	201
//	@Failure	400	{object}	response.ResFailed
//	@Failure	404	{object}	response.ResFailed
//	@Failure	500	{object}	response.ResFailed
//	@Router		/jobs [post]
func (h *jobHandler) Create(c *gin.Context) {
	req := &model.CreateJobReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseFailed(err.Error()))
		return
	}

	if err := h.jobUsecase.Create(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseFailed(err.Error()))
		return
	}

	c.Status(http.StatusCreated)
}

// FindJob godoc
//
//	@Summary	find job
//	@Tags		jobs
//	@Produce	json
//	@Param		page		query		int		false	"Page number"
//	@Param		size		query		int		false	"Page size"
//	@Param		keyword		query		string	false	"Keyword for job title or description"
//	@Param		companyName	query		string	false	"Company name"
//
//	@Success	200			{array}		response.ResSuccess
//	@Failure	400			{object}	response.ResFailed
//	@Failure	404			{object}	response.ResFailed
//	@Failure	500			{object}	response.ResFailed
//	@Router		/jobs [get]
func (h *jobHandler) FindJob(c *gin.Context) {
	keyword := c.Query("keyword")
	companyName := c.Query("companyName")
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 50
	}

	query := make(map[string]interface{})
	query["keyword"] = keyword
	query["companyName"] = companyName
	query["page"] = page
	query["size"] = size

	jobs, meta, err := h.jobUsecase.FindJob(c, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseFailed(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.ResponseSuccess(jobs, meta))
}
