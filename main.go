package main

import (
	"encoding/json"
	"net/http"
)

type HandlerData struct {
	Word   string `json:"word"`
	Number int64  `json:"number"`
}

func makeHandlerData(word string, num int64) *HandlerData {
	return &HandlerData{
		Word:   word,
		Number: num,
	}
}

func stepHandler(writer http.ResponseWriter, _ *http.Request) {
	data := []*HandlerData{
		makeHandlerData("first", 4),
		makeHandlerData("second", 8),
		makeHandlerData("third", 13),
	}
	encodedData, _ := json.Marshal(data)

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")

	if _, err := writer.Write(encodedData); err != nil {
		panic(err)
	}
}

func rootHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(404)

	if _, err := writer.Write([]byte("Just try '/step'")); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/step", stepHandler)

	if err := http.ListenAndServe("localhost:3001", nil); err != nil {
		panic(err)
	}
}
