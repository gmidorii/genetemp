package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
	"strings"
)

type Class struct {
	Name string
	Path string
}

func main() {
	// read param
	param, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var class Class
	err = yaml.Unmarshal(param, &class)
	classMap := map[string]string{
					"[className]": class.Name,
					"[dir]": class.Path}

	// read template
	temp, err := os.Open("template/service.java")
	if err != nil {
		log.Fatal(err)
	}
	defer temp.Close()

	re, _ := regexp.Compile("\\[.*?\\]")
	sc := bufio.NewScanner(temp)
	for i := 0; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			break
		}
		text := sc.Text()
		match := re.FindAllStringSubmatch(text, -1)
		for key, value := range classMap {
			for _, matchValue := range match {
				for _, m := range matchValue {
					if key == m {
						text = strings.Replace(text, m, value, -1)
					}
				}
			}
		}
		fmt.Println(text)
	}

	// output file
	err = ioutil.WriteFile("output/"+class.Name, []byte(class.Name), 0755)
	if err != nil {
		log.Fatal(err)
	}
}
