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

type ArtistController struct {
	ArtistService services.ArtistService
}

func NewArtistController(artistService services.ArtistService) *ArtistController {
	return &ArtistController{
		ArtistService: artistService,
	}
}

func (controller *ArtistController) GetArtistByUserId(ctx *gin.Context) {
	userId := ctx.Query("id")

	res, err := controller.ArtistService.GetArtistByUserId(userId)
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

func (controller *ArtistController) GetArtistByArtistId(ctx *gin.Context) {
	artistId := ctx.Query("id")

	res, err := controller.ArtistService.GetArtistByArtistId(artistId)
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

func (controller *ArtistController) CreateArtist(ctx *gin.Context) {
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
	userId := ctx.PostForm("userId")

	//_, err = controller.ArtistService.GetArtistByUserId(userId)
	//if err == nil {
	//	webResponse := response.WebResponse{
	//		Code:    http.StatusBadRequest,
	//		Message: errors.New("verification is requested").Error(),
	//		Data:    nil,
	//	}
	//
	//	ctx.Header("Content-Type", "application/json")
	//	ctx.JSON(http.StatusBadRequest, webResponse)
	//	return
	//}

	//_, err = controller.ArtistService.GetUnverifiedArtistByArtistId(userId)
	//if err == nil {
	//	webResponse := response.WebResponse{
	//		Code:    http.StatusBadRequest,
	//		Message: errors.New("you are verified").Error(),
	//		Data:    nil,
	//	}
	//
	//	ctx.Header("Content-Type", "application/json")
	//	ctx.JSON(http.StatusBadRequest, webResponse)
	//	return
	//}

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
	artistRequest := request.ArtistRequest{
		UserId:      userId,
		Description: &description,
		Banner:      &imageUrl,
	}
	err = controller.ArtistService.CreateArtist(artistRequest)
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

func (controller *ArtistController) UpdateVerifyArtist(ctx *gin.Context) {
	artistId := ctx.Query("id")

	err := controller.ArtistService.UpdateVerifyArtist(artistId)
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

func (controller *ArtistController) GetUnverifiedArtist(ctx *gin.Context) {
	res, err := controller.ArtistService.GetUnverifiedArtist()
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

func (controller *ArtistController) DeleteArtist(ctx *gin.Context) {
	artistId := ctx.Query("id")
	userId := ctx.Query("userId")
	fmt.Println(userId)

	err := controller.ArtistService.DeleteArtist(userId, artistId)
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
