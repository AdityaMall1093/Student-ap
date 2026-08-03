package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Adityamall1093/student-api/internal/config"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to student api"))
	})

	//setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	err := server.ListenAndServe()

	fmt.Println("server started")
	if err != nil {
		log.Fatal("Failed to start server")
	}
	fmt.Println("server")

}
