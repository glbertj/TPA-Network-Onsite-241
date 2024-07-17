package controller

import (
	"back-end/data/request"
	"back-end/data/response"
	"back-end/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PlayController struct {
	PlayService services.PlayService
}

func NewPlayController(playService services.PlayService) *PlayController {
	return &PlayController{PlayService: playService}
}

func (c *PlayController) Create(ctx *gin.Context) {
	var req request.PlayRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		webResponse := response.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	err = c.PlayService.Create(req)
	if err != nil {
		webResponse := response.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (c *PlayController) Get8LastPlayedSongByUser(ctx *gin.Context) {

	userId := ctx.Query("id")

	result, err := c.PlayService.GetLastPlayedSongByUser(userId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    result,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *PlayController) GetLastPlayedSongByUser(ctx *gin.Context) {

	userId := ctx.Query("id")

	result, err := c.PlayService.GetLastPlayedSongByUser(userId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    result,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
