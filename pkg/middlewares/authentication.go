package middlewares

import (
	"fmt"
	"net/http"

	"../logger"
	"../token"
)

const noTokenMessage = "NT Bad Request"
const invalidTokenMessage = "Unauthorized request"

//AuthenticationMiddleware logs requests
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := token.ExtractToken(r)
		if tokenString == "" {
			logger.Log("No Bearer Token In Authorization Header", logger.INFO)
			http.Error(w, noTokenMessage, http.StatusBadRequest)
			return
		}
		isOk, _ := token.VerifyToken(tokenString)

		if isOk == false {
			logger.Log(fmt.Sprintf("%v %v", invalidTokenMessage, tokenString), logger.INFO)
			http.Error(w, invalidTokenMessage, http.StatusUnauthorized)
			return
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
