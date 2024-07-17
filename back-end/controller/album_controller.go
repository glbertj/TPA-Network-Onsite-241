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

type AlbumController struct {
	AlbumService services.AlbumService
}

func NewAlbumController(albumService services.AlbumService) *AlbumController {
	return &AlbumController{AlbumService: albumService}
}

func (a AlbumController) GetAlbumByArtist(ctx *gin.Context) {
	artistId := ctx.Query("id")
	res, err := a.AlbumService.GetAlbumsByArtist(artistId)
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

func (a AlbumController) GetRandomAlbum(ctx *gin.Context) {
	res, err := a.AlbumService.GetRandomAlbum()
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

func (a AlbumController) CreateAlbum(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	title := ctx.PostForm("title")
	artistId := ctx.PostForm("artistId")
	typeAlbum := ctx.PostForm("type")
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
	album := request.AlbumRequest{
		Title:    title,
		Type:     typeAlbum,
		Banner:   imageUrl,
		ArtistId: artistId,
	}
	res, err := a.AlbumService.CreateAlbum(album)
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
