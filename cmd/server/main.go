package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"todolist-api/config"
	"todolist-api/routes"
)

func main() {
	config.InitDatabase()

	router := httprouter.New()

	router.GET("/health", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"message":"Server is running"}`))
	})

	routes.Init(router, config.GetDB())

	port := ":" + config.GetServerPort()
	log.Println("Server Started on " + port)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}