package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type homeResponse struct {
	Hello string `json:"hello"`
}

const port string = ":8081"

func buildJenkins(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["JIRA_ISSUE_KEY"]
	fmt.Println(keys[0])
}

func home(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(homeResponse{
		Hello: "Hello docker!",
	})
}

func main() {
	fmt.Printf("Server listening on http://127.0.0.1%s/\n", port)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/build/jenkins", buildJenkins)
	log.Fatal(http.ListenAndServe(port, router))
}
