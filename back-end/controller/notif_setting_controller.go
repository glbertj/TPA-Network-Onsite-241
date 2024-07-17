package controller

import (
	"back-end/data/request"
	"back-end/data/response"
	"back-end/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NotificationSettingController struct {
	NotificationSettingService services.NotificationSettingService
}

func NewNotificationSettingController(notificationSettingService services.NotificationSettingService) *NotificationSettingController {
	return &NotificationSettingController{NotificationSettingService: notificationSettingService}
}

func (n *NotificationSettingController) CreateNotificationSetting(userId string) error {
	return n.NotificationSettingService.CreateNotificationSetting(userId)
}

func (n *NotificationSettingController) GetSettingBySettingId(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := n.NotificationSettingService.GetSettingBySettingId(id)
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

func (n *NotificationSettingController) UpdateSetting(ctx *gin.Context) {
	req := request.NotificationSettingRequest{}
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

	err = n.NotificationSettingService.UpdateNotificationSetting(req)
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
