package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

type serverconfig struct {
	Port         string `yaml:"Port"`
	ReadTimeout  int    `yaml:"ReadTimeout"`
	WriteTimeout int    `yaml:"WriteTimeout"`
}

func config() *serverconfig {
	filename, _ := filepath.Abs("./config.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	serverconfig := &serverconfig{}
	err = yaml.Unmarshal(yamlFile, &serverconfig)
	if err != nil {
		panic(err)
	}
	return serverconfig
}

var (
	Port         = config().Port
	ReadTimeout  = time.Duration(config().ReadTimeout) * time.Second
	WriteTimeout = time.Duration(config().WriteTimeout) * time.Second
)

func server() {
	handler := http.NewServeMux()

	handler.HandleFunc("/hello/", helloHandler)
	s := http.Server{
		Addr:         Port,
		Handler:      handler,
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
	}
	log.Fatal(s.ListenAndServe())

}
func main() {
	server()
	// Route handles & endpoints

}

type resp struct {
	Message string `json:"Message"`
	Error   string `json:"Error"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := resp{
		Message: "hello",
		Error:   "null",
	}
	respJs, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Write(respJs)
}
