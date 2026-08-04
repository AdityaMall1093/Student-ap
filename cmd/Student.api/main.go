package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// student "github.com/Adityamall1093/student-api/internal/http/handlers"
	// "github.com/Adityamall1093/student-api/internal/config"
	"github.com/Adityamall1093/student-api/internal/config"
	student "github.com/Adityamall1093/student-api/internal/http/handlers"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()

	// router.HandleFunc("GET/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome to student api"))
	// })
	router.HandleFunc("POST /api/student", student.New())

	//setup server

	// server := http.Server{
	// 	Addr:    cfg.Addr,
	// 	Handler: router,
	// }

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	slog.Info("server started ", slog.String("address", cfg.Addr))

	fmt.Printf("server started %s", cfg.Addr)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT)

	go func() {

		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	fmt.Println("server")
	<-done

	slog.Info("shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")

}
