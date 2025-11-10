package main

import (
	"booksmanage/config"
	"booksmanage/server"
	"booksmanage/sql"
	"fmt"
	"net/http"
)

func main() {
	config.LoadConfig()
	sql.InitDB()
	http.HandleFunc("/login", server.Login)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("met error", err)
	}
}
