package controller

import (
	"back-end/data/response"
	"back-end/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type AdvertisementController struct {
	AdvertisementService services.AdvertisementService
}

func NewAdvertisementController(advertisementService services.AdvertisementService) *AdvertisementController {
	return &AdvertisementController{AdvertisementService: advertisementService}
}

func (a AdvertisementController) GetRandomAdvertisement(ctx *gin.Context) {
	res, err := a.AdvertisementService.GetRandomAdvertisement()
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

func (a AdvertisementController) StreamAdv(c *gin.Context) {
	id := c.Query("id")
	adv, err := a.AdvertisementService.GetAdvertisementById(id)
	if err != nil {
		webResponse := response.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	file, err := os.Open(adv.Link)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error opening file")
		return
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error getting file info")
		return
	}
	fileSize := fileInfo.Size()
	// Assume 128 kbps bitrate, 10 seconds of audio
	chunkSize := int64(128 * 1024 * 10 / 8)
	rangeHeader := c.GetHeader("Range")
	var start, end int64
	if rangeHeader != "" {
		_, err := fmt.Sscanf(rangeHeader, "bytes=%d-", &start)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid range header")
			return
		}
	}
	end = start + chunkSize
	if end > fileSize {
		end = fileSize
	}
	c.Header("Content-Type", "audio/mpeg")
	c.Header("Accept-Ranges", "bytes")
	c.Header("Content-Length", fmt.Sprintf("%d", end-start))
	c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end-1, fileSize))
	c.Status(http.StatusPartialContent)
	_, err = file.Seek(start, 0)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error seeking file")
		return
	}
	_, err = io.CopyN(c.Writer, file, end-start)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error copying file")
		return
	}
}
