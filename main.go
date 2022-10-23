package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	code := getStatusCode(r)
	writeResponse(code, w)
}

func getStatusCode(request *http.Request) int {
	statusInt, _ := strconv.Atoi(request.URL.Query().Get("code"))
	return statusInt
}

func writeResponse(statusCode int, response http.ResponseWriter) {

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)

	fileName := strconv.Itoa(statusCode) + ".json"
	jsonFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo: "+fileName, err)
	}

	response.Write(jsonFile)
}
