package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/golang/gddo/httputil/header"
)

//MalformedRequest type
type MalformedRequest struct {
	IsSuccess bool
	Err       string
	Value     allObject
}

func (mr *MalformedRequest) Error() string {
	return mr.Err
}

type allObject interface {
}

//DecodeJSONBody decodes JSON Body
func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst allObject) error {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			return &MalformedRequest{IsSuccess: false, Err: msg}
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return &MalformedRequest{IsSuccess: false, Err: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			return &MalformedRequest{IsSuccess: false, Err: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &MalformedRequest{IsSuccess: false, Err: msg}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &MalformedRequest{IsSuccess: false, Err: msg}

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &MalformedRequest{IsSuccess: false, Err: msg}

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return &MalformedRequest{IsSuccess: false, Err: msg}

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		return &MalformedRequest{IsSuccess: false, Err: msg}
	}

	return nil
}

//EncodeJSONResponse encodes JSON response
func EncodeJSONResponse(w http.ResponseWriter, err error, dst allObject) {
	response := MalformedRequest{}

	response.IsSuccess = true
	if err != nil {
		response.IsSuccess = false
		response.Err = err.Error()
	}

	if dst != nil {
		response.Value = dst
	}
	json.NewEncoder(w).Encode(response)
}
