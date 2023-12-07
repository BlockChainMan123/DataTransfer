package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BuildResp(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Code:    "0",
		Message: "ok",
		Data:    data,
	}
	bytes, err := json.Marshal(response);
	if err != nil {
		BuildErr(w, "json encode error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func BuildErr(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Code:    "1",
		Message: message,
		Data:    map[string]int{},
	}
	bytes, err := json.Marshal(response);
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"code": "1", "message": "json encode error."}`)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write(bytes)
}

func BuildRet(w http.ResponseWriter, result string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func handlePageSizeAndStart(r *http.Request) (int32, int) {
	// read page_size and start from url, refer to the site:
	// https://golangcode.com/get-a-url-parameter-from-a-request/
	page_size, ok := r.URL.Query()["page_size"]
	start, ok2 := r.URL.Query()["start"]

	PageSize := int64(10)
	Start := int64(1)

	if ok && (len(page_size) >= 1) {
		temp, err := strconv.ParseInt(page_size[0], 10, 32)
		if err == nil {
			PageSize = temp
		}
	}
	if ok2 && (len(start) >= 1) {
		temp, err := strconv.ParseInt(start[0], 10, 64)
		if err == nil {
			Start = temp
		}
	}

	// ensure PageSize bigger than 1
	if PageSize < 1 {
		PageSize = 1
	}
	// ensure Start bigger than 1
	if Start < 1 {
		Start = 1
	}
	return int32(PageSize), int(Start)
}
