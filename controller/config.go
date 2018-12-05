package controller

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
	"log"
	"fmt"
)

var Filename = "./config.yml"

type Serverconfig struct {
	Port         string `yaml:"Port"`
	ReadTimeout  int    `yaml:"ReadTimeout"`
	WriteTimeout int    `yaml:"WriteTimeout"`
}

	func (s *Serverconfig) validate() func() {
	    fmt.Println("Parsing Config",Filename)
	    return func() {
		if s.Port == "" {
			log.Fatal("Config Pars ERROR Port Required ", s.Port)
		}
		if s.ReadTimeout == 0 {
			log.Fatal("Config Pars ERROR ReadTimeout Requierd", s.ReadTimeout)
		}
		if s.WriteTimeout == 0 {
			log.Fatal("Config Pars ERROR WriteTimeout Requierd", s.WriteTimeout)
		}
	}
}

func ServerConfig() *Serverconfig {
	Filename, _ := filepath.Abs(Filename)
	yamlFile, err := ioutil.ReadFile(Filename)
	if err != nil {
		panic(err)
	}
	Serverconfig := Serverconfig{}
	err = yaml.Unmarshal(yamlFile, &Serverconfig)
	f := Serverconfig.validate()
	f()
	fmt.Println("Start server with parametrs\nServer Port" + Serverconfig.Port)
	if err != nil {
		log.Fatal(err)
	}
	return &Serverconfig
}
