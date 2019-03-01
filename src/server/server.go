package main

import (
	dbClient "github.com/Thai56/simple-grpc/src/db"
	mux "github.com/gorilla/mux"
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
	server := &APIServer{}
	r := mux.NewRouter()

	// Paths
	r.HandleFunc("/", server.Pong)

	listenPort := ":9090"
	server.http = http.Server{
		Addr:         listenPort,
		Handler:      r,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	// create a grpc connection to a database
	dbStruct := dbClient.Database{}
	db, err := dbStruct.Serve()

	if err != nil {
		log.Errorf("Failed to serve the client db: %v", err)
	}

	log.Infof("Successfully served the db client, %v", err)

	pong, err := db.DbConn.Ping().Result()
	if err != nil {
		log.Errorf("Failed to connect to db:%v", err)
		return
	}
	log.Infof("ping:%s", pong)

	log.Infof("Listening on port %s", listenPort)

	err = server.http.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to on %s: %s", listenPort, err)
		return
	}

	// create middleware that will be used to process preflight headers to allow json requests to the front end
}

//Pong - will respond to a ping
func (s *APIServer) Pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	log.Info("Hitting Pong!")
}
