package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	rs "rce.amopdev/m/v2/pkg/service"
	reg "rce.amopdev/m/v2/util/regex"
)

var (
	service = rs.NewRemoteService("", "")
)

// Create api that writes code from request to the code/ folder for different languages
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing request...")
	// Read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	code := string(body)

	service.SetSettings("python", "")
    err, resp := service.RunCode(code)
	service.RemoveFile()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	jData, err := json.Marshal(reg.RegReplace(resp))
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jData)
}

func main() {
    http.HandleFunc("/run", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}