package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/leonmaia/vod-api/api"
	"github.com/leonmaia/vod-api/persistence"
)

var (
	dbName     = flag.String("database", "transmissions", "database name")
	dbUser     = flag.String("username", "root", "database user name")
	dbPassword = flag.String("password", "root", "database password")

	httpPort = flag.String("http", "8000", "http port")
)

func createRoutes(tHandler *api.TransmissionHandler) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/transmissions/{id}", tHandler.GetURL).Methods("GET")
	r.HandleFunc("/transmissions", tHandler.Create).Methods("POST")
	r.HandleFunc("/health", api.HealthCheckHandler)

	return r
}

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", *dbUser, *dbPassword, *dbName))
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxIdleConns(50)
	defer db.Close()

	handler := api.TransmissionHandler{Repository: persistence.Repository{DB: db}}

	r := createRoutes(&handler)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("127.0.0.1:%s", *httpPort),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Println("Listening on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
