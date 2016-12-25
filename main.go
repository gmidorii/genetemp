package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type Class struct {
	Name      string
	Path      string
	Extension string
	Template  string
}

func main() {
	// set flag
	c := flag.String("c", "config", "loading config file path")
	flag.Parse()
	configReader, err := os.Open(*c)
	errorCheck(err)
	defer configReader.Close()

	for n, class := range readClasses(configReader) {
		classMap := map[string]string{
			"[name]": class.Name,
			"[path]": class.Path}

		dir, err := os.Getwd()
		errorCheck(err)

		for _, path := range strings.Split(class.Path, "/") {
			dir = filepath.Join(dir, path)
		}
		classMap["[path]"] = strings.Replace(class.Path, "/", ".", -1) + "." + class.Name

		// make directory
		if !dirExist(dir) {
			err := os.MkdirAll(dir, os.ModePerm)
			errorCheck(err)
		}

		// create output file and writer
		file := filepath.Join(dir, class.Name+class.Extension)
		fp, err := os.Create(file)
		errorCheck(err)
		writer := bufio.NewWriter(fp)

		// read template file
		temp, err := os.Open(class.Template)
		if err != nil {
			log.Fatal(err)
		}
		sc := bufio.NewScanner(temp)
		reg, _ := regexp.Compile("\\[.*?\\]")
		for sc.Scan() {
			if err := sc.Err(); err != nil {
				break
			}

			text := sc.Text()
			match := reg.FindAllString(text, -1)
			if len(match) == 0 {
				writeFile(text, writer)
				continue
			}

			for key, value := range classMap {
				for _, m := range match {
					if key == m {
						text = strings.Replace(text, m, value, -1)
					}
				}
			}
			writeFile(text, writer)
		}

		fmt.Println("--------------------------------")
		fmt.Println("no: " + strconv.Itoa(n))
		fmt.Println("template: " + class.Template)
		fmt.Println("config: " + *c)
		fmt.Println("create: " + file)
		fmt.Println("--------------------------------")
	}
}

func dirExist(dirname string) bool {
	dir, err := os.Stat(dirname)
	if err != nil {
		return false
	}
	return dir.IsDir()
}

func writeFile(text string, writer *bufio.Writer) {
	writer.Write([]byte(text + "\n"))
	writer.Flush()
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readClasses(r io.Reader) []Class {
	param, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	var classes []Class
	err = yaml.Unmarshal(param, &classes)
	if err != nil {
		log.Fatal(err)
	}
	return classes
}
