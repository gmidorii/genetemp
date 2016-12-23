package main

import (
	"bufio"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
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

	// read param
	param, err := ioutil.ReadFile(*c)
	if err != nil {
		log.Fatal(err)
	}
	var classes []Class
	err = yaml.Unmarshal(param, &classes)

	for n, class := range classes {
		classMap := map[string]string{
			"[name]": class.Name,
			"[path]": class.Path}

		// get current directory
		dir, err := os.Getwd()
		errorCheck(err)

		var pack string
		paths := strings.Split(class.Path, "/")
		for _, path := range paths {
			dir = filepath.Join(dir, path)
			pack = pack + "." + path
		}
		pack = pack + "." + class.Name
		pack = strings.TrimLeft(pack, ".")
		classMap["[path]"] = pack

		// make directory
		if !dirExist(dir) {
			err := os.MkdirAll(dir, os.ModePerm)
			errorCheck(err)
		}

		file := filepath.Join(dir, class.Name+class.Extension)
		fp, err := os.Create(file)
		errorCheck(err)
		writer := bufio.NewWriter(fp)

		// read template
		temp, err := os.Open(class.Template)
		if err != nil {
			log.Fatal(err)
		}
		reg, _ := regexp.Compile("\\[.*?\\]")
		sc := bufio.NewScanner(temp)
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
