package route

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64  `json:omitempty`
	UserName string `json:omitempty`
	Priv     bool   `json:omitempty`
	Password []byte `json:omitempty`
}

func Register(responseW http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	// logic part of log in
	fmt.Println("username:", request.Form["username"])
	fmt.Println("password:", request.Form["password"])
	fmt.Println("password1:", request.Form["password1"])
	if request.PostForm.Get("password") != request.PostForm.Get("password1") {
		http.Redirect(responseW, request, "/register/", 302)

	} else {
		username := request.PostForm.Get("username")
		password := request.PostForm.Get("password")
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(username, string(hash))
	}

}
