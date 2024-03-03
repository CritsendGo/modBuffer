package modBuffer

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

var folderTest = "/tmp/test/"
var size = 2
var entryExample any = time.Now()
var tesContainerType = reflect.TypeOf(TestContainer{})

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
type TestContainer struct {
	Val  string
	Date time.Time
}

func TestAll(t *testing.T) {
	Debug = false
	Log = true
	err := os.RemoveAll(folderTest)
	if err != nil {
		t.Fatal("Unable to clean folder buffer", err)
	}
	bu, err := NewBuffer(folderTest, size)
	if err != nil {
		t.Fatal("Unable to create buffer", err)
	}

	// Need to bet 1 buffer and 0 in folder
	rep := addContent(bu, "First")
	if rep != "First" {
		t.Fatal("Add One Entry , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err, "CONTENT", rep)
	}

	// Need to bet 0 buffer and 0 in folder
	rep = getContent(bu)
	if rep != "First" {
		t.Fatal("NEED TO BE BUFFER=0 AND FOLDER=0 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err, "CONTENT", rep)
	}

	// Need to bet 1 buffer and 0 in folder
	rep = addContent(bu, "Second")
	if rep != "Second" {
		t.Fatal("Add One Entry , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err, "CONTENT", rep)
	}

	// Need to bet 1 buffer and 0 in folder
	rep = addContent(bu, "Third")
	if rep != "Third" {
		t.Fatal("Add One Entry , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err, "CONTENT", rep)
	}

	// Need to bet 1 buffer and 0 in folder
	rep = addContent(bu, "Fourth")
	if rep != "Fourth" {
		t.Fatal("Add One Entry , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err, "CONTENT", rep)
	}

	rep = getContent(bu)
	if rep != "Second" {
		t.Fatal("NEED TO BE BUFFER=0 AND FOLDER=0 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err, "CONTENT", rep)
	}
	rep = getContent(bu)
	if rep != "Third" {
		t.Fatal("NEED TO BE BUFFER=1 AND FOLDER=1 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err, "CONTENT", rep)
	}
	rep = addContent(bu, "Fifth")
	if rep != "Fifth" {
		t.Fatal("Add One Entry , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err, "CONTENT", rep)
	}
	// Reading before next loop of folder buffer
	rep = getContent(bu)
	if rep != "Fifth" {
		t.Fatal("NEED TO BE BUFFER=1 AND FOLDER=1 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err, "CONTENT", rep)
	}
	time.Sleep(2 * ReadingInterval)
	rep = getContent(bu)
	if rep != "Fourth" {
		t.Fatal("NEED TO BE BUFFER=1 AND FOLDER=1 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err, "CONTENT", rep)
	}
}

func addContent(bu *CSBuffer, val string) string {

	obj := TestContainer{val, time.Now()}
	err := bu.Add(obj)
	if Log == true {
		fmt.Println("ADDING CONTENT")
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err, "CONTENT", obj.Val)
	}

	return obj.Val
}
func getContent(bu *CSBuffer) string {

	obj := TestContainer{}
	val, err := bu.Get()
	if err != nil {
		fmt.Println("GETTING CONTENT BUT EMPTY")
		return ""
	}
	d, _ := json.Marshal(val)
	_ = json.Unmarshal(d, &obj)
	if Log == true {
		fmt.Println("GETTING CONTENT")
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err, "CONTENT", obj.Val)
	}
	return obj.Val
}
