package api

import (
	"database/sql"
	"net/http"
	"strconv"

	e "github.com/bobbybof/timesheet/error"
	h "github.com/bobbybof/timesheet/internal/helper"
	"github.com/bobbybof/timesheet/internal/repository"
	"github.com/gin-gonic/gin"
)

type getUserRequest struct {
	Id string `uri:"id" binding:"required,numeric,gt=0"`
}

type updateUserRequest struct {
	Name string `json:"name" binding:"required"`
	Rate int32  `json:"rate" binding:"required"`
}

func (server *Server) GetUser(ctx *gin.Context) {
	var reqUri getUserRequest

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, e.ErrorHttpResponse(err, ""))
		return
	}

	uid, err := strconv.ParseInt(reqUri.Id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, e.ErrorHttpResponse(err, ""))
		return
	}

	user, err := server.store.GetUser(ctx, int32(uid))

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, h.SuccessHttpResponse("User not found", nil))
			return
		}
		ctx.JSON(http.StatusInternalServerError, e.ErrorHttpResponse(err, ""))
		return
	}

	ctx.JSON(http.StatusOK, h.SuccessHttpResponse("success", user))
}

func (server *Server) UpdateUser(ctx *gin.Context) {
	var reqBody updateUserRequest
	var reqUri getUserRequest

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, e.ErrorHttpResponse(err, ""))
		return
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, e.ErrorHttpResponse(err, ""))
		return
	}

	uid, err := strconv.ParseInt(reqUri.Id, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, e.ErrorHttpResponse(err, ""))
		return
	}

	user, err := server.store.UpdateUser(ctx, repository.UpdateUserParams{
		Name: reqBody.Name,
		Rate: reqBody.Rate,
		ID:   int32(uid),
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, e.ErrorHttpResponse(err, ""))
		return
	}

	ctx.JSON(http.StatusOK, h.SuccessHttpResponse("success", user))
}
