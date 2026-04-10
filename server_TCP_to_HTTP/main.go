package main

import (
	"log"
	"net/http"
)

func main() {
	api := &api{
		addr: ":8080",
	}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr: api.addr,
		Handler :mux ,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsershandler)

	log.Println("Server chay tai cong 8080")
	if err := srv.ListenAndServe(); err != nil{
		panic(err)
	}
}