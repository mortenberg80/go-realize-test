package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type HelloWorldResponse struct {
	UserName string `json:"user_name"`
}

// HelloWorld
func HelloWorld(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	jsonSession, err := json.Marshal(CreateHelloWorldResponse("Foo"))
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Fprintln(w, string(jsonSession))
	return
}

func CreateHelloWorldResponse(name string) *HelloWorldResponse {
	response := HelloWorldResponse{
		UserName: name,
	}

	return &response
}
