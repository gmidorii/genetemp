package main

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Class struct {
	Name	string
	Path	string
}

func main()  {
	config, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var class Class
	err = yaml.Unmarshal(config, &class)

	err = ioutil.WriteFile("output/" + class.Name, []byte(class.Name), 0755)
	if err != nil {
		log.Fatal(err)
	}
}