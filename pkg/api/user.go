package api

import (
	"errors"
	"net/http"

	"../logger"
	"../token"
	"../util"
)

//TokenResponse is response of Login Api
type TokenResponse struct {
	Token string `json:"access_token"`
}

//UserInfo holds user info
type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const usernameCannotBeEmpty = "Username cannot be empty"
const passwordCannotBeEmpty = "Password cannot be empty"
const userPassWrong = "Username or Password is wrong"

//Login Login
func Login(w http.ResponseWriter, r *http.Request) {

	userInfo := UserInfo{}

	err := util.DecodeJSONBody(w, r, &userInfo)

	if err != nil {
		logger.Log(err.Error(), logger.ERROR)
		util.EncodeJSONResponse(w, err, &userInfo)
		return
	}

	if userInfo.Username == "" {
		err = errors.New(usernameCannotBeEmpty)
		util.EncodeJSONResponse(w, err, nil)
		return
	}

	if userInfo.Password == "" {
		err = errors.New(passwordCannotBeEmpty)
		util.EncodeJSONResponse(w, err, nil)
		return
	}

	userInfoMap := map[string]string{}

	userInfoMap["user_name"] = userInfo.Username
	userInfoMap["is_authorized"] = "true"

	var tokenString string
	tokenString, err = token.CreateToken(userInfoMap)

	response := TokenResponse{
		Token: tokenString,
	}

	util.EncodeJSONResponse(w, err, response)

}
