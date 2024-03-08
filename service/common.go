package service

import (
	"encoding/json"
	"net/http"
)

func doResponse(w http.ResponseWriter, herr *HttpError, result any) {
	resp := ToHttpResult(result, herr)
	bts, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Write(bts)
}
func ToHttpResult(result any, herr *HttpError) HttpResult {
	if herr != nil {
		return HttpResult{Result: result, Error: herr}
	} else {
		return HttpResult{Result: result, Error: nil}
	}
}
func NewHttpError(code ErrorCode, msg string) *HttpError {
	return &HttpError{Code: string(code), Message: msg}
}

type HttpResult struct {
	Result any        `json:"result"`
	Error  *HttpError `json:"error,omitempty"`
}
type HttpError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
type ErrorCode string
