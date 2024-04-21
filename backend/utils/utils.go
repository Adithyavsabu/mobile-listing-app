package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// decodes a JSON request into the provided interface (mainly any reference to struct) and handles errors internally
func DecodeJsonRequest(w http.ResponseWriter, r *http.Request, data interface{}) {
	err := json.NewDecoder(r.Body).Decode(&data)
	fmt.Println("i am in decodejson request and the value decoded:", data)
	if err != nil {
		fmt.Println("Couldnt fetch data", err)
		return
	}
}

func FetchId(w http.ResponseWriter, r *http.Request) int {
	params := mux.Vars(r)
	productID, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("error fetching Id", err)
		return 0
	}
	return productID
}
