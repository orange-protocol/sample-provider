package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type SampleDPRequest struct {
	Address string `json:"address"`
}

type SampleDPResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

func SampleGetBodyDP(w http.ResponseWriter, r *http.Request) {

	key := r.Header.Get("x-api-key")
	if !strings.EqualFold(key, "test") {
		doResponse(w, NewHttpError(INVALID_PARAM, "Unauthorized"), nil)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		doResponse(w, NewHttpError(INVALID_PARAM, err.Error()), nil)
		return
	}
	param := SampleDPRequest{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		doResponse(w, NewHttpError(INVALID_PARAM, err.Error()), nil)
		return
	}
	doResponse(w, nil, &SampleDPResponse{Address: param.Address, Balance: "20"})
	return

}

func SampleGetGetUrlDP(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("x-api-key")
	fmt.Printf("key: %v\n", key)
	if !strings.EqualFold(key, "test") {
		doResponse(w, NewHttpError(INVALID_PARAM, "Unauthorized"), nil)
		return
	}
	addr := r.URL.Query().Get("address")
	chain := r.URL.Query().Get("chain")
	fmt.Printf("addr:%s,chain:%s\n", addr, chain)
	doResponse(w, nil, &SampleDPResponse{Address: addr, Balance: "30"})

}
