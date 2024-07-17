package controller

import (
	"back-end/data/response"
	"back-end/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchController struct {
	SearchService services.SearchService
}

func NewSearchController(searchService services.SearchService) *SearchController {
	return &SearchController{SearchService: searchService}
}

func (c SearchController) Search(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	res, err := c.SearchService.Search(keyword)
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
