package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	defer fmt.Println("Exited!")

	port := getEnv("PORT", "8993")
	msg := fmt.Sprintf("Listening on port %s ...\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, msg)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}
