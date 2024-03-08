package service

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var (
	INVALID_PARAM  ErrorCode = "INVALID_PARAM"
	INTERNAL_ERROR ErrorCode = "INTERNAL_ERROR"
	NOT_IMPLEMENT  ErrorCode = "NOT_IMPLEMENT"
)

type SampleAPRequest struct {
	Result struct {
		Address string `json:"address"`
		Balance string `json:"balance"`
	} `json:"result"`
}
type SampleAPResponse struct {
	Score string `json:"score"`
}

func SampleAP(w http.ResponseWriter, r *http.Request) {

	key := r.Header.Get("x-api-key")
	// fmt.Printf("key: %v\n", key)
	if !strings.EqualFold(key, "test") {
		doResponse(w, NewHttpError(INVALID_PARAM, "Unauthorized"), nil)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		doResponse(w, NewHttpError(INVALID_PARAM, err.Error()), nil)
		return
	}
	param := SampleAPRequest{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		doResponse(w, NewHttpError(INVALID_PARAM, err.Error()), nil)
		return
	}

	balance, err := strconv.ParseFloat(param.Result.Balance, 64)
	if err != nil {
		doResponse(w, NewHttpError(INVALID_PARAM, err.Error()), nil)
		return
	}
	score := "0"
	if balance <= 100 {
		score = "10"
	} else {
		score = "30"
	}
	doResponse(w, nil, &SampleAPResponse{Score: score})

}
