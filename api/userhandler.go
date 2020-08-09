package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Handle(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("Handler called")
	handle("go-service")
	response, _ := json.Marshal("Hello from go-service")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func handle(msg string) string {
	fmt.Println("User handler called......... ", msg)
	return "Hello"+msg
}


