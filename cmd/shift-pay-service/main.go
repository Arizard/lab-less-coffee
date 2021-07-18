package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"

	handlers "github.com/arizard/lab-less-coffee/cmd/shift-pay-service/gorilla-handlers"
)

func getEnv(name string, fallback string) string {
	v := os.Getenv(name)
	if v == "" {
		return fallback
	}
	return v
}

func main() {
	port := getEnv("PORT", "8080")

	r := mux.NewRouter()
	r.HandleFunc("/ping", handlers.PingHandler)
	r.HandleFunc("/ping", handlers.PingHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		fmt.Println(err)
	}

}
