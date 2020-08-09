package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Handle(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("Handler called")
	HandleRequest("go-service")
	response, _ := json.Marshal("Hello from go-service")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func HandleRequest(msg string) string {
	fmt.Println("User handler called......... ", msg)
	return "Hello"+msg
}


