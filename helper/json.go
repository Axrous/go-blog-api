package helper

import (
	"encoding/json"
	"net/http"
)

func ResponToBody(writer http.ResponseWriter, result interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(result)
	PanicIfError(err)
}