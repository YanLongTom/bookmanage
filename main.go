package main

import (
	bookserver "booksmanage/server"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/login", bookserver.Login)
	http.HandleFunc("/logout", bookserver.Logout)
	http.HandleFunc("/view", bookserver.View)
	http.HandleFunc("/addbook", bookserver.AddBook)
	http.HandleFunc("/deletebook", bookserver.DeleteBook)
	http.HandleFunc("/changebook", bookserver.ChangeBook)
	http.HandleFunc("/readebook", bookserver.ReadeBook)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("have error", err)
	}
}
