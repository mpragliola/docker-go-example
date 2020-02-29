package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	initializeService()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello world")
	})

	http.Handle("/", r)

	var serverPort = os.Getenv("SERVICE_INTERNAL_PORT")

	if serverPort == "" {
		serverPort = "80"
	}

	log.Println(fmt.Sprintf("Starting server, listening [port %s]", serverPort))
	log.Fatal(http.ListenAndServe(":"+serverPort, nil))
}

func initializeService() {
	log.Println(fmt.Sprintf("Service started [PID %d]", os.Getpid()))

	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file, using env variables as they are")
	}
}
