package server

import (
	"booksmanage/model"
	"booksmanage/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	u1 := model.User{}
	if r.Method != "POST" {
		w.Write([]byte("method invalid"))
		return
	}
	ret, err := io.ReadAll(r.Body)
	err = json.Unmarshal(ret, &u1)
	if err != nil {
		w.Write([]byte("input invalid"))
	} else {
		result := saveUser(u1)
		w.WriteHeader(200)
		w.Write([]byte("hello " + u1.Name + "\n"))
		w.Write([]byte(result))
	}
}
func saveUser(user model.User) string {
	err, ret := sql.SaveUser(&user)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return ret
}
