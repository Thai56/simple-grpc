package main

import (
	mux "github.com/gorilla/mux"
	dbClient "github.com/simple-grpc/src/db"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// APIServer - is the bridge from REST requests to GRPC requests
type APIServer struct {
	http     http.Server
	dbClient interface{}
}

func main() {
	s := &APIServer{}
	r := mux.NewRouter()

	// Paths
	r.HandleFunc("/", s.Pong)

	listenPort := ":9090"
	s.http = http.Server{
		Addr:         listenPort,
		Handler:      r,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	log.Infof("Listening on port %s", listenPort)

	err := s.http.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to on %s: %s", listenPort, err)
		return
	}

	// create a grpc connection to a database
	err = dbClient.Serve()

	if err != nil {
		log.Errorf("Failed to serve the client db: %v", err)
	}

	log.Infof("Successfully served the db client, %v")

	// create middleware that will be used to process preflight headers to allow json requests to the front end
}

//Pong - will respond to a ping
func (s *APIServer) Pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	log.Info("Hitting Pong!")
}
