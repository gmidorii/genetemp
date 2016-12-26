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

	"reflect"

	"gopkg.in/yaml.v2"
)

var version = "1.0"

type Class struct {
	Name      string
	ClassName string
	Path      string
	Package   string
	Extension string
	Template  string
}

func main() {
	// set flag
	var v bool
	flag.BoolVar(&v, "v", false, "show version")
	var c string
	flag.StringVar(&c, "c", "config", "loading config file path")
	flag.Parse()

	if v {
		fmt.Println("version: ", version)
		return
	}

	configReader, err := os.Open(c)
	if err != nil {
		log.Fatal(err)
	}
	defer configReader.Close()

	for n, class := range ReadClasses(configReader) {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		for _, path := range strings.Split(class.Path, "/") {
			dir = filepath.Join(dir, path)
		}

		class.Path = strings.Replace(class.Path, "/", ".", -1)

		// make directory
		if !dirExist(dir) {
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}

		// create output file and writer
		file := filepath.Join(dir, class.ClassName+class.Extension)
		fp, err := os.Create(file)
		if err != nil {
			log.Fatal(err)
		}
		writer := bufio.NewWriter(fp)

		// read template file
		temp, err := os.Open(class.Template)
		if err != nil {
			log.Fatal(err)
		}
		sc := bufio.NewScanner(temp)
		reg, _ := regexp.Compile("\\[.*?\\]")

		classMap := ConvertToMap(class)
		for sc.Scan() {
			if err := sc.Err(); err != nil {
				break
			}

			text := sc.Text()
			match := reg.FindAllString(text, -1)
			if len(match) == 0 {
				WriteFile(text, writer)
				continue
			}

			for key, value := range classMap {
				for _, m := range match {
					if key == m {
						text = strings.Replace(text, m, value, -1)
					}
				}
			}
			WriteFile(text, writer)
		}

		fmt.Println("--------------------------------")
		fmt.Println("no: " + strconv.Itoa(n))
		fmt.Println("template: " + class.Template)
		fmt.Println("config: " + c)
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

func WriteFile(text string, writer *bufio.Writer) {
	writer.Write([]byte(text + "\n"))
	writer.Flush()
}

func ReadClasses(r io.Reader) []Class {
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

func ConvertToMap(class Class) map[string]string {
	rt := reflect.TypeOf(class)
	rv := reflect.ValueOf(class)
	classMap := map[string]string{}
	for i := 0; i < rt.NumField(); i++ {
		key := "[" + strings.ToLower(rt.Field(i).Name) + "]"
		classMap[key] = rv.Field(i).Interface().(string)
	}
	return classMap
}
