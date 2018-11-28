package main

import (
	"gobots/controller"
	"net/http"
	"time"
	"log"
	"html/template"
	"fmt"
)

func main() {
	 Serverconfig := controller.ServerConfig()

	handler := http.NewServeMux()
    handler.HandleFunc("/login/" ,login)
	s := http.Server{
		 Addr: Serverconfig.Port,
		 Handler: handler,
		 ReadTimeout: time.Duration(Serverconfig.ReadTimeout) * time.Second,
		 WriteTimeout: time.Duration(Serverconfig.WriteTimeout) * time.Second,
	}
	    log.Fatal(s.ListenAndServe())
	}

func login(responseW http.ResponseWriter, requset *http.Request) {
	t,err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(responseW,err.Error(),400)
		return
	}
	if err := t.Execute(responseW, nil); err != nil {
		fmt.Println(err)
	    return
	}
}

