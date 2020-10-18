package api

import (
	"net/http"

	"../util"
)

// VerifyTest is test method for verifing token
func VerifyTest(w http.ResponseWriter, r *http.Request) {

	util.EncodeJSONResponse(w, nil, "Verify Ok")
	return
}
