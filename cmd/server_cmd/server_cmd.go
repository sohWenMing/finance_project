package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	database_connection "github.com/sohWenMing/finance_project/internal/database"
	"github.com/sohWenMing/finance_project/internal/server"
)

func main() {

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
	db, err := database_connection.Connect_db()
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}
	defer db.Close()
	server := server.InitServer(db)
	<-quitChan
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	err = server.Shutdown(ctx)
	if err != nil {
		log.Printf("error shutting down server: %v", err)
		return
	}
	fmt.Println("server shutdown gracefully")
}
