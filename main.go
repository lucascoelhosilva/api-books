package main

import (
	"api-books/pkg/routers"
	"net/http"
	"time"
)

func main()  {
	router := routers.InitRouter()
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}
	srv.ListenAndServe()
}