package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"http_server/model"
)

const (
	EXAMPLE_HTTP_ADDRESS = ":8080"
	ERR_MSG              = "\"retcode\":500, \"errmsg\": \"Internal Error!\""
)

func main()  {
	startHttpServer()
}
func startHttpServer() {
	fmt.Println("Start HTTP server on http://localhost:8080...")
	http.HandleFunc("/user/", handleApi1)
	http.HandleFunc("/order/", handleApi2)
	http.ListenAndServe(EXAMPLE_HTTP_ADDRESS, nil)
}
func handleApi1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API-1 receive: ", r.RequestURI)
	apiRequest := model.ApiRequest{}
	apiRequest.Retcode = 200

	user := model.User{}
	user.Id = "user10086"
	user.Name = "user-1"

	apiRequest.Res = user
	apiReqByte, err := json.Marshal(apiRequest)

	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	if err != nil {
		w.Write([]byte(ERR_MSG))
	} else {
		w.Write(apiReqByte)
	}
}

func handleApi2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API-2 receive: ", r.RequestURI)
	apiRequest := model.ApiRequest{}
	apiRequest.Retcode = 200

	order := model.Order{}
	order.Id = "order10000"
	order.Title = "order-1"
	apiRequest.Res = order

	apiReqByte, err := json.Marshal(apiRequest)
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	if err != nil {
		w.Write([]byte(ERR_MSG))
	} else {
		w.Write(apiReqByte)
	}
}
