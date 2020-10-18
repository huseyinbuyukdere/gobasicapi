package token

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"../logger"
	"github.com/dgrijalva/jwt-go"
)

//CreateToken creates token
func CreateToken(dataMap map[string]string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}

	for key, value := range dataMap {
		atClaims[key] = value
	}

	var timeout int
	timeout, err = strconv.Atoi(os.Getenv("TOKEN_TIMEOUT"))
	if err != nil {
		logger.Log(err.Error(), logger.WARNING)
		timeout = 15
	}

	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(timeout)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

//VerifyToken verifies token
func VerifyToken(token string) (bool, error) {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

//ExtractToken extracts token from request
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
