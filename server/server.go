package main

import (
	//	"encoding/json"
	//"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	//	"github.com/gorilla/handlers"
	rhandlers "github.com/remotejob/real_estate_redux_go/server/handlers"
	"net/http"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	})
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api", rhandlers.AllRecords)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8080",c.Handler(gorillaRoute))
}
