package main

import (
	"bufio"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"flag"
)

type Class struct {
	Name string
	Path string
}

func main() {
	c := flag.String("c", "config", "loading config file path")
	t := flag.String("t", "template", "loading template file path")
	flag.Parse()

	// read param
	param, err := ioutil.ReadFile(*c)
	if err != nil {
		log.Fatal(err)
	}
	var class Class
	err = yaml.Unmarshal(param, &class)
	classMap := map[string]string{
		"[className]": class.Name,
		"[dir]":       class.Path}

	// read template
	temp, err := os.Open(*t)
	if err != nil {
		log.Fatal(err)
	}
	defer temp.Close()

	fp, err := os.OpenFile("output/" + class.Name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(fp)

	reg, _ := regexp.Compile("\\[.*?\\]")
	sc := bufio.NewScanner(temp)
	for i := 0; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			break
		}

		text := sc.Text()
		match := reg.FindAllString(text, -1)
		if match == nil {
			writer.WriteString("\n")
			continue
		}

		for key, value := range classMap {
			for _, m := range match {
				if key == m {
					text = strings.Replace(text, m, value, -1)
				}
			}
		}
		writer.WriteString(text)
		writer.WriteString("\n")
		writer.Flush()
	}
}
