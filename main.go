package main

import (
	"ginauth101/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRoute()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("ListenAndServer::", err)
	}
}
