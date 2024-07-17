package database

import (
	"back-end/config"
	"back-end/data/request"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"net/http"
)

type GoogleConf struct {
	GoogleConfig *oauth2.Config
}

func NewGoogle(cnf *config.Config) *GoogleConf {
	return &GoogleConf{
		GoogleConfig: &oauth2.Config{
			ClientID:     cnf.Google.ClientID,
			ClientSecret: cnf.Google.ClientSecret,
			Endpoint:     google.Endpoint,
			RedirectURL:  cnf.Google.RedirectURL,
			Scopes:       []string{"profile", "email"},
		},
	}
}

//func (g *Google) googleCallback(ctx *gin.Context) {
//	code := ctx.Query("code")
//
//	token, err := g.googleConfig.Exchange(context.Background(), code)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"error": err.Error(),
//		})
//		return
//	}
//
//	userInfo, err := GetUserInfo(token.AccessToken)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"error": err.Error(),
//		})
//		return
//	}
//
//}

func GetUserInfo(accessToken string) (request.GoogleRequest, error) {
	userEndPoint := "https://www.googleapis.com/oauth2/v2/userinfo"
	resp, err := http.Get(fmt.Sprintf("%s?access_token=%s", userEndPoint, accessToken))
	if err != nil {
		return request.GoogleRequest{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return request.GoogleRequest{}, err
	}
	googleResp := request.GoogleRequest{
		Username: userInfo["name"].(string),
		Email:    userInfo["email"].(string),
		GoogleId: userInfo["id"].(string),
	}
	return googleResp, nil
}
