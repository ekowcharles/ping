package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	defer fmt.Println("Exited!")

	version, err := ioutil.ReadFile(".version")
	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, rq *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, fmt.Sprintf("pong %s", string(version)))
	})

	log.Fatal(http.ListenAndServe(":8993", nil))
}
