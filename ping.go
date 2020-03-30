package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	defer fmt.Println("Exited!")

	port := getEnv("PORT", "8993")

	v, err := ioutil.ReadFile(".version")
	if err != nil {
		panic(err.Error())
	}

	version := string(v)

	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, fmt.Sprintf("Version: %s, Port: %s", version, port))
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, rq *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, fmt.Sprintf("pong %s", version))
	})

	fmt.Printf("Listening on port %s ...\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}
