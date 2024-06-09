package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	e "github.com/bobbybof/timesheet/error"
	h "github.com/bobbybof/timesheet/internal/helper"
	"github.com/bobbybof/timesheet/internal/repository"
	"github.com/gin-gonic/gin"
)

type createActivityRequest struct {
	Name      string `json:"name" binding:"required"`
	DateStart string `json:"date_start" binding:"required"`
	DateEnd   string `json:"date_end" binding:"required"`
	ProjectID int    `json:"project_id" binding:"required"`
	UserID    int    `json:"user_id" binding:"required"`
}

type updateActivityRequest struct {
	Name      string `json:"name" binding:"required"`
	DateStart string `json:"date_start" binding:"required"`
	DateEnd   string `json:"date_end" binding:"required"`
	ProjectID int    `json:"project_id" binding:"required"`
}

type getActivityByIdRequest struct {
	Id string `uri:"id" binding:"required"`
}

type getDataQueryParams struct {
	Search  string `form:"search"`
	Project string `form:"project"`
	UserId  string `form:"user_id" binding:"required"`
}

func (server *Server) CreateActivity(ctx *gin.Context) {
	var req createActivityRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, e.ErrorHttpResponse(err, ""))
		return
	}

	dateStart, err := h.DateFormat(req.DateStart)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, e.ErrorHttpResponse(err, "invalid date format"))
		return
	}
	dateEnd, err := h.DateFormat(req.DateEnd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, e.ErrorHttpResponse(err, "invalid date format"))
		return
	}
	if dateStart.After(dateEnd) {
		ctx.JSON(http.StatusBadRequest, e.ErrorHttpResponse(errors.New("invalid date"), "date start must before date end"))
		return
	}

	activity, err := server.store.CreateActivity(ctx, repository.CreateActivityParams{
		Name:      req.Name,
		DateStart: dateStart,
		DateEnd:   dateEnd,
		ProjectID: int32(req.ProjectID),
		UserID:    int32(req.UserID),
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, e.ErrorHttpResponse(err, "Something went wrong while try to create activity"))
		return
	}

	ctx.JSON(http.StatusOK, activity)
}

func (server *Server) UpdateActivity(ctx *gin.Context) {
	var reqBody updateActivityRequest
	var reqUri getActivityByIdRequest

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, e.ErrorHttpResponse(err, ""))
		return
	}

	activityId, err := strconv.ParseInt(reqUri.Id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, e.ErrorHttpResponse(err, ""))
		return
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, e.ErrorHttpResponse(err, ""))
		return
	}

	dateStart, err := h.DateFormat(reqBody.DateStart)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, e.ErrorHttpResponse(err, "invalid date format"))
		return
	}
	dateEnd, err := h.DateFormat(reqBody.DateEnd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, e.ErrorHttpResponse(err, "invalid date format"))
		return
	}
	if dateStart.After(dateEnd) {
		ctx.JSON(http.StatusBadRequest, e.ErrorHttpResponse(errors.New("invalid date"), "date start must before date end"))
		return
	}

	activity, err := server.store.UpdateActivity(ctx, repository.UpdateActivityParams{
		Name:      reqBody.Name,
		DateStart: dateStart,
		DateEnd:   dateEnd,
		ProjectID: int32(reqBody.ProjectID),
		ID:        int32(activityId),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, e.ErrorHttpResponse(err, "Something went wrong while try to create activity"))
		return
	}

	ctx.JSON(http.StatusOK, activity)
}

func (server *Server) DeleteActivity(ctx *gin.Context) {
	var reqUri getActivityByIdRequest

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, e.ErrorHttpResponse(err, ""))
		return
	}

	activityId, err := strconv.ParseInt(reqUri.Id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, e.ErrorHttpResponse(err, ""))
		return
	}

	err = server.store.DeleteActivity(ctx, int32(activityId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, e.ErrorHttpResponse(err, ""))
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})

}

func (server *Server) GetUserActivities(ctx *gin.Context) {
	var reqQuery getDataQueryParams
	var projects []string

	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, e.ErrorHttpResponse(err, ""))
		return
	}

	if reqQuery.Project != "" {
		projects = strings.Split(reqQuery.Project, ",")
		for i, project := range projects {
			projects[i] = strings.TrimSpace(project)
		}
	}

	uid, err := strconv.ParseInt(reqQuery.UserId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, e.ErrorHttpResponse(err, ""))
		return
	}

	activities, err := server.store.GetUserActivites(ctx, projects, reqQuery.Search, int32(uid))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, e.ErrorHttpResponse(err, ""))
		return
	}

	ctx.JSON(http.StatusOK, h.SuccessHttpResponse("success", activities))
}
