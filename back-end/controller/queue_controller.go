package controller

import (
	"back-end/data/request"
	"back-end/data/response"
	"back-end/model"
	"back-end/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type QueueController struct {
	QueueService services.QueueService
}

func NewQueueController(queueService services.QueueService) *QueueController {
	return &QueueController{QueueService: queueService}
}

func (q QueueController) ClearQueue(ctx *gin.Context) {
	key := ctx.Query("key")
	err := q.QueueService.ClearQueue(key)
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
		Data:    "cleared",
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (q QueueController) Enqueue(ctx *gin.Context) {
	key := ctx.Query("key")
	var song request.QueueRequest
	err := ctx.ShouldBindJSON(&song)
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

	songs := model.Song{
		SongId:      song.SongId,
		Title:       song.Title,
		ArtistId:    song.ArtistId,
		AlbumId:     song.AlbumId,
		ReleaseDate: song.ReleaseDate,
		Duration:    song.Duration,
		File:        song.File,
		Album:       song.Album,
		Play:        song.Play,
		Artist:      song.Artist,
	}
	err = q.QueueService.Enqueue(key, songs)
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

func (q QueueController) Dequeue(ctx *gin.Context) {
	key := ctx.Query("key")
	song, err := q.QueueService.Dequeue(key)
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
	res := response.SongResponse{
		SongId:      song.SongId,
		Title:       song.Title,
		ArtistId:    song.ArtistId,
		AlbumId:     song.AlbumId,
		ReleaseDate: song.ReleaseDate,
		Duration:    song.Duration,
		File:        song.File,
		Album:       song.Album,
		Play:        song.Play,
		Artist:      song.Artist,
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    res,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (q QueueController) GetQueue(ctx *gin.Context) {
	key := ctx.Query("key")
	song, err := q.QueueService.GetQueue(key)
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
	res := response.SongResponse{
		SongId:      song.SongId,
		Title:       song.Title,
		ArtistId:    song.ArtistId,
		AlbumId:     song.AlbumId,
		ReleaseDate: song.ReleaseDate,
		Duration:    song.Duration,
		File:        song.File,
		Album:       song.Album,
		Play:        song.Play,
		Artist:      song.Artist,
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    res,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (q QueueController) GetAllQueue(ctx *gin.Context) {
	key := ctx.Query("key")
	songs, err := q.QueueService.GetAllQueue(key)
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
		Data:    songs,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (q QueueController) RemoveFromQueue(ctx *gin.Context) {
	key := ctx.Query("key")
	index := ctx.Query("index")
	idx, err := strconv.Atoi(index)
	if err != nil {
		webResponse := response.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    "Failed to convert index to integer",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	err = q.QueueService.RemoveFromQueue(key, idx)
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
		Data:    "Success",
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
