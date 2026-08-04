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

	// router.HandleFunc("GET/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome to student api"))
	// })
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Student API"))
	})

	//setup server

	// server := http.Server{
	// 	Addr:    cfg.Addr,
	// 	Handler: router,
	// }

	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	fmt.Printf("server started %s", cfg.HTTPServer.Addr)

	go func() {

		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	fmt.Println("server")

}
