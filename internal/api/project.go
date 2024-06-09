package api

import (
	"net/http"

	e "github.com/bobbybof/timesheet/error"
	h "github.com/bobbybof/timesheet/internal/helper"
	"github.com/gin-gonic/gin"
)

type createProjectRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) CreateProject(ctx *gin.Context) {
	var reqBody createProjectRequest

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, e.ErrorHttpResponse(err, ""))
		return
	}

	proj, err := server.store.CreateProject(ctx, reqBody.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, e.ErrorHttpResponse(err, ""))
		return
	}

	ctx.JSON(http.StatusOK, h.SuccessHttpResponse("success", proj))
}

func (server *Server) GetProjects(ctx *gin.Context) {
	projects, err := server.store.GetProjects(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, e.ErrorHttpResponse(err, ""))
		return
	}

	ctx.JSON(http.StatusOK, h.SuccessHttpResponse("success", projects))
}
