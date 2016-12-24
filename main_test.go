package main

import (
	"bufio"
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestReadClasses(t *testing.T) {
	s := "- name: Test1\n  path: output/test1\n  extension: .java\n  template: template/test1.java\n"
	stdin := bytes.NewBufferString(s)
	r := io.Reader(stdin)
	classes := readClasses(r)
	if len(classes) != 1 {
		t.Fatalf("expected read class number 1: actual %d", len(classes))
	}

	expected := Class{"Test1", "output/test1", ".java", "template/test1.java"}
	actual := classes[0]
	if !reflect.DeepEqual(actual, expected) {
		t.Error("struct is not equal")
	}
}

func TestWriteFile(t *testing.T) {
	stdout := new(bytes.Buffer)
	w := bufio.NewWriter(stdout)
	expected := "hoge"
	writeFile(expected, w)

	if bytes.Compare([]byte(expected+"\n"), stdout.Bytes()) != 0 {
		t.Error("write string not match")
	}
}
