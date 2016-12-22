package main

import (
	"bufio"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"fmt"
)

type Class struct {
	Name string
	Path string
	Extension string
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

	// get current directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	output := filepath.Join(dir, "output")
	// make directory
	if !dirExist(output) {
		err := os.Mkdir(output, os.ModePerm)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	file := filepath.Join(output, class.Name + class.Extension)
	fp, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(fp)

	reg, _ := regexp.Compile("\\[.*?\\]")
	sc := bufio.NewScanner(temp)
	for sc.Scan() {
		if err := sc.Err(); err != nil {
			break
		}

		text := sc.Text()
		match := reg.FindAllString(text, -1)
		if len(match) == 0 {
			writer.WriteString(text)
			writer.WriteString("\n")
			writer.Flush()
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

	fmt.Println("--------------------------------")
	fmt.Println("template: " + *t)
	fmt.Println("config: " + *c)
	fmt.Println("create: " + file)
	fmt.Println("--------------------------------")
}

func dirExist(dirname string) bool {
	dir, err := os.Stat(dirname)
	if err != nil {
		return false
	}
	return dir.IsDir()
}
