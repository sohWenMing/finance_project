package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

func InitServer(*sql.DB) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Port not defined in environment, defaulting to port :8080")
		port = ":8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler)
	server := &http.Server{
		Addr:    "0.0.0.0" + port,
		Handler: mux,
	}
	fmt.Printf("server is listening on port %s\n", port)
	go func() {

		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("error operating server: %v", err)
		}
	}()
	return server
}
