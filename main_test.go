package main

import (
	"bufio"
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestReadClasses(t *testing.T) {
	s := "- name: Test1\n  classname: Test1Service\n  path: output/test1\n  package: output.test1\n  extension: .java\n  template: template/test1.java\n"

	stdin := bytes.NewBufferString(s)
	r := io.Reader(stdin)
	classes := ReadClasses(r)
	if len(classes) != 1 {
		t.Fatalf("expected read class number 1: actual %d", len(classes))
	}

	expected := Class{"Test1", "Test1Service", "output/test1", "output.test1", ".java", "template/test1.java"}
	actual := classes[0]
	if !reflect.DeepEqual(actual, expected) {
		t.Error("struct is not equal")
	}
}

func TestWriteFile(t *testing.T) {
	stdout := new(bytes.Buffer)
	w := bufio.NewWriter(stdout)
	expected := "hoge"
	WriteFile(expected, w)

	if bytes.Compare([]byte(expected+"\n"), stdout.Bytes()) != 0 {
		t.Error("write string not match")
	}
}

func TestConvertToMap(t *testing.T) {
	class := Class{
		Name:      "name",
		ClassName: "classname",
		Path:      "path",
		Package:   "package",
		Extension: "ext",
		Template:  "tmp"}
	m := ConvertToMap(class)

	if v, ok := m["[name]"]; !ok || v != "name" {
		t.Errorf("map value incorrect {ok: %t, v:%s}", ok, v)
	}
	if v, ok := m["[classname]"]; !ok || v != "classname" {
		t.Errorf("map value incorrect {ok: %t, v:%s}", ok, v)
	}
	if v, ok := m["[path]"]; !ok || v != "path" {
		t.Errorf("map value incorrect {ok: %t, v:%s}", ok, v)
	}
	if v, ok := m["[package]"]; !ok || v != "package" {
		t.Errorf("map value incorrect {ok: %t, v:%s}", ok, v)
	}
	if v, ok := m["[extension]"]; !ok || v != "ext" {
		t.Errorf("map value incorrect {ok: %t, v:%s}", ok, v)
	}
	if v, ok := m["[template]"]; !ok || v != "tmp" {
		t.Errorf("map value incorrect {ok: %t, v:%s}", ok, v)
	}
}
