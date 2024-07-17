package controller

import (
	"back-end/data/request"
	"back-end/data/response"
	"back-end/database"
	"back-end/services"
	"back-end/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type UserController struct {
	UserService services.UserService
	g           *database.GoogleConf
}

func NewUserController(userService services.UserService, g *database.GoogleConf) *UserController {
	u := UserController{UserService: userService, g: g}
	return &u

}

func (u *UserController) GoogleCallback(ctx *gin.Context) {
	code := ctx.Query("code")

	token, err := u.g.GoogleConfig.Exchange(context.Background(), code)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userInfo, err := database.GetUserInfo(token.AccessToken)
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

	user, err := u.UserService.LoginWithGoogle(userInfo)
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

	fmt.Println("token : " + user.Token)
	ctx.SetCookie("jwt", user.Token, 24*60*60, "/", "", false, true)

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    user,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (u *UserController) UpdateVerificationStatus(ctx *gin.Context) {
	id := ctx.Query("id")
	err := u.UserService.UpdateVerificationStatus(id)
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

func (u *UserController) Authenticate(ctx *gin.Context) {
	req := request.AuthRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {

		webResponse := response.WebResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
			Data:    nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusUnauthorized, webResponse)
		return

	}

	user, err := u.UserService.Authenticate(req)
	if err != nil {
		webResponse := response.WebResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
			Data:    nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusUnauthorized, webResponse)
		return
	}
	ctx.SetCookie("jwt", user.Token, int(24*time.Hour.Seconds()), "/", "", false, true)

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    user,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (u *UserController) GetCurrentUser(ctx *gin.Context) {
	token, err := ctx.Cookie("jwt")
	if err != nil {
		webResponse := response.WebResponse{
			Code:    http.StatusUnauthorized,
			Message: "Cookie not Found",
			Data:    nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusUnauthorized, webResponse)
		return
	}

	user, err := u.UserService.GetCurrentUser(token)
	if err != nil {
		webResponse := response.WebResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
			Data:    nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusUnauthorized, webResponse)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    user,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (u *UserController) Register(ctx *gin.Context) {
	req := request.UserRequest{}
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

	user, err := u.UserService.Register(req)
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
	link := "http://localhost:5173/verify-email?id=" + user.VerifyLink
	err = utils.SendEmail(user.Email, "Verification Email", fmt.Sprintf("Please click this link to verify your email: <a href=\"%s\">Verify Email</a> ", link))
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
		Data:    user,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (u *UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Query("id")
	user, err := u.UserService.GetUserById(id)
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
		Data:    user,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (u *UserController) UpdateUserProfile(ctx *gin.Context) {
	req := request.UserUpdateRequest{}
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

	user, err := u.UserService.UpdateUserProfile(req)
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
		Data:    user,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (u *UserController) GetAllUser(ctx *gin.Context) {
	users, err := u.UserService.GetAllUser()
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
		Data:    users,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (u *UserController) SignOut(ctx *gin.Context) {
	ctx.SetCookie("jwt", "", -1, "/", "", false, true)
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    "Sign Out Success",
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (u *UserController) Forgot(ctx *gin.Context) {
	email := ctx.Query("email")
	user, err := u.UserService.ForgotPassword(email)
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
	link := "http://localhost:5173/reset-pass?id=" + user.VerifyLink
	err = utils.SendEmail(user.Email, "Change Password", fmt.Sprintf("Please click this link to change your password: <a href=\"%s\">Chang Password</a> ", link))

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

func (u *UserController) GetUserByVerifyLink(ctx *gin.Context) {
	link := ctx.Query("id")
	user, err := u.UserService.GetUserByVerifyLink(link)
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
		Data:    user,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (u *UserController) ResetPassword(ctx *gin.Context) {
	req := request.ResetPasswordRequest{}
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

	err = u.UserService.ChangePassword(req.Password, req.UserId)
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

func (u *UserController) Logout(ctx *gin.Context) {
	err := u.UserService.Logout(ctx.Query("id"))
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
	ctx.SetCookie("jwt", "", -1, "/", "", false, true)
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    "Success",
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (u *UserController) UpdateProfilePicture(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	artistId := ctx.PostForm("userId")
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
	err = u.UserService.UpdateProfilePicture(artistId, imageUrl)
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
