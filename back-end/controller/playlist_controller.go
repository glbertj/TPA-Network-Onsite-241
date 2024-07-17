package controller

import (
	"back-end/data/request"
	"back-end/data/response"
	"back-end/services"
	"back-end/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type PlaylistController struct {
	PlaylistService services.PlaylistService
}

func NewPlaylistController(playlistService services.PlaylistService) *PlaylistController {
	return &PlaylistController{PlaylistService: playlistService}
}

func (p *PlaylistController) GetPlaylistByUserId(ctx *gin.Context) {
	id := ctx.Query("id")

	res, err := p.PlaylistService.GetByUserID(id)
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
		Data:    res,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (p *PlaylistController) GetPlaylistById(ctx *gin.Context) {
	id := ctx.Query("id")

	res, err := p.PlaylistService.GetPlaylistByID(id)
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
		Data:    res,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (p *PlaylistController) CreateDetail(ctx *gin.Context) {
	var req request.PlayListDetailRequest
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

	err = p.PlaylistService.CreateDetail(req)
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

func (p *PlaylistController) DeletePlaylistDetail(ctx *gin.Context) {
	id := ctx.Query("id")
	userId := ctx.Query("userId")
	detailId := ctx.Query("detId")

	err := p.PlaylistService.DeletePlaylistDetailByID(userId, id, detailId)
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

func (p *PlaylistController) DeletePlaylist(ctx *gin.Context) {
	id := ctx.Query("id")
	userId := ctx.Query("userId")
	err := p.PlaylistService.DeletePlaylistByID(userId, id)
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

func (p *PlaylistController) CreatePlaylist(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
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
	description := ctx.PostForm("description")
	title := ctx.PostForm("title")
	userId := ctx.PostForm("userId")
	filename := strings.Replace(utils.GenerateUUID(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	image := fmt.Sprintf("%s.%s", filename, fileExt)
	err = ctx.SaveUploadedFile(file, fmt.Sprintf("./assets/images/%s", image))
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
	imageUrl := fmt.Sprintf("http://localhost:4000/public/images/%s", image)
	playlistRequest := request.PlayListRequest{
		UserID:      userId,
		Title:       title,
		Description: description,
		Image:       imageUrl,
	}
	err = p.PlaylistService.Create(playlistRequest)
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
