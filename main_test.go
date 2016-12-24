package main

import (
	"os"
	"testing"
	"reflect"
)

func TestReadClasses(t *testing.T) {
	dir, _ := os.Getwd()
	configReader, err := os.Open(dir + "/config/test.yaml")
	if err != nil {
		t.Error(err)
	}

	classes := readClasses(configReader)
	if len(classes) != 2 {
		t.Errorf("expected read class number 2: actual %d", len(classes))
	}

	expected := Class{"Test1", "output/test1", ".java", "template/test1.java"}
	actual := classes[0]
	if !reflect.DeepEqual(actual, expected)  {
		t.Error("struct is not equal")
	}
}
