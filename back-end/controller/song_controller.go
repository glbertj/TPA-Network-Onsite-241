package controller

import (
	"back-end/data/request"
	"back-end/data/response"
	"back-end/services"
	"back-end/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type SongController struct {
	SongService services.SongService
}

func NewSongController(songService services.SongService) *SongController {
	return &SongController{SongService: songService}
}

func (s SongController) GetAllSong(ctx *gin.Context) {
	res, err := s.SongService.GetAllSong()
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

func (s SongController) GetSongById(ctx *gin.Context) {
	id := ctx.Query("id")

	res, err := s.SongService.GetSongById(id)
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

func (s SongController) GetSongByArtist(ctx *gin.Context) {
	artistId := ctx.Query("id")

	res, err := s.SongService.GetSongByArtist(artistId)
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

func (s SongController) GetSongByAlbum(ctx *gin.Context) {
	albumId := ctx.Query("id")

	res, err := s.SongService.GetSongByAlbum(albumId)
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

func (s SongController) CreateSong(ctx *gin.Context) {
	file, err := ctx.FormFile("song")
	title := ctx.PostForm("title")
	albumId := ctx.PostForm("albumId")
	artistId := ctx.PostForm("artistId")
	duration := ctx.PostForm("duration")
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
	songs := fmt.Sprintf("%s.%s", filename, "mp3")
	err = ctx.SaveUploadedFile(file, fmt.Sprintf("./assets/songs/%s", songs))
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
	songUrl := fmt.Sprintf("./assets/songs/%s", songs)

	durationInt, err := strconv.Atoi(duration)
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
	song := request.SongRequest{
		SongId:      utils.GenerateUUID(),
		Title:       title,
		ArtistId:    artistId,
		AlbumId:     albumId,
		ReleaseDate: time.Now(),
		Duration:    durationInt,
		File:        songUrl,
	}
	err = s.SongService.CreateSong(song)
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

func (s SongController) TestMusic(c *gin.Context) {
	songId := c.Query("id")
	music, err := s.SongService.GetSongById(songId)
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
	file, err := os.Open(music.File)
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

func (s SongController) StreamMusic(c *gin.Context) {
	songId := c.Query("id")
	music, err := s.SongService.GetSongById(songId)
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
	file, err := os.Open(music.File)
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
	chunkSize := int64(1024 * 200)
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

//func (s SongController) StreamMusic(c *gin.Context) {
//	songId := c.Query("id")
//	music, err := s.SongService.GetSongById(songId)
//	if err != nil {
//		webResponse := response.WebResponse{
//			Code:    http.StatusBadRequest,
//			Message: err.Error(),
//			Data:    nil,
//		}
//
//		c.Header("Content-Type", "application/json")
//		c.JSON(http.StatusBadRequest, webResponse)
//		return
//	}
//
//	file, err := os.Open(music.File)
//	if err != nil {
//		panic(err)
//	}
//	defer func(file *os.File) {
//		err := file.Close()
//		if err != nil {
//
//		}
//	}(file)
//
//	c.Header("Content-Type", "audio/mpeg")
//	c.Header("Transfer-Encoding", "chunked")
//	//c.Header("Content-Encoding", "gzip")
//	//c.Header("Content-Type", "text/html")
//
//	buffer := make([]byte, 1024)
//	writer := c.Writer
//	for {
//		n, err := file.Read(buffer)
//		if err != nil && err != io.EOF {
//			c.String(http.StatusInternalServerError, "Error reading file.")
//			return
//		}
//		if n == 0 {
//			break
//		}
//		if _, err := writer.Write(buffer[:n]); err != nil {
//			c.String(http.StatusInternalServerError, "Error writing chunk.")
//			return
//		}
//		writer.(http.Flusher).Flush() // Flush the buffer to the client
//	}
//}
