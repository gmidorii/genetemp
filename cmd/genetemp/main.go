package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Class struct {
	Name	string
	Path	string
}

func main()  {
	buf, err := ioutil.ReadFile("template/template.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var class Class
	err = yaml.Unmarshal(buf, &class)
	fmt.Println(class.Name)
}