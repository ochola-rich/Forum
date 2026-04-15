package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("this is the home page\n"))
	if err != nil {
		slog.Error("error writing the response")
	}
}



func ping(w http.ResponseWriter, r *http.Request){
	params := r.URL.Query()

	userList := params["user"]

	var output bytes.Buffer

	user := "user"
	output.WriteString("Hello ")
	if len(userList) > 0 {
		user = userList[0]
	}
	output.WriteString(user)
	output.WriteString("!\n")

	_,err := w.Write(output.Bytes())
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", root)
	mux.HandleFunc("/ping", ping)

	fmt.Println("server running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
