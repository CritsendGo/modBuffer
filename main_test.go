package modBuffer

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
)

var folderTest = "/tmp/test/"
var size = 2
var entryExample any = time.Now()

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.

func TestCreateBuffer(t *testing.T) {
	Debug = false
	_, err := NewBuffer(folderTest, size)
	if err != nil {
		t.Fatal("Unable to create buffer", err)
	}
}
func TestAll(t *testing.T) {
	Debug = true
	err := os.RemoveAll(folderTest)
	if err != nil {
		t.Fatal("Unable to clean folder buffer", err)
	}
	bu, err := NewBuffer(folderTest, size)
	time.Sleep(100 * time.Millisecond)
	if err != nil {
		t.Fatal("Unable to create buffer", err)
	}

	fmt.Println("Adding Item")
	err = bu.Add(&entryExample)
	time.Sleep(100 * time.Millisecond)
	// Need to bet 1 buffer and 0 in folder
	if len(bu.data) != 1 || bu.SizeNew() != 0 {
		t.Fatal("Add One Entry , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}

	fmt.Println("Getting Item")
	_, err = bu.Get()
	time.Sleep(100 * time.Millisecond)
	// Need to bet 0 buffer and 0 in folder
	if len(bu.data) != 0 || bu.SizeNew() != 0 {
		t.Fatal("NEED TO BE BUFFER=0 AND FOLDER=0 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}

	fmt.Println("Adding Item")
	err = bu.Add(&entryExample)
	time.Sleep(100 * time.Millisecond)
	if len(bu.data) != 1 || bu.SizeNew() != 0 {
		t.Fatal("NEED TO BE BUFFER=0 AND FOLDER=0 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}

	fmt.Println("Adding Item")
	err = bu.Add(&entryExample)
	time.Sleep(100 * time.Millisecond)
	if len(bu.data) != 2 || bu.SizeNew() != 0 {
		t.Fatal("NEED TO BE BUFFER=0 AND FOLDER=0 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}

	fmt.Println("Adding Item")
	err = bu.Add(&entryExample)
	time.Sleep(100 * time.Millisecond)
	if len(bu.data) != 2 || bu.SizeNew() != 1 {
		t.Fatal("NEED TO BE BUFFER=0 AND FOLDER=0 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}

	fmt.Println("Getting Item")
	_, err = bu.Get()
	time.Sleep(100 * time.Millisecond)
	if len(bu.data) != 1 || bu.SizeNew() != 1 {
		t.Fatal("NEED TO BE BUFFER=1 AND FOLDER=1 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}

	fmt.Println("Reading Item")
	bu.ScanFolder()
	if len(bu.data) != 2 || bu.SizeNew() != 0 {
		t.Fatal("NEED TO BE BUFFER=2 AND FOLDER=0 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}

	fmt.Println("Getting Item")
	_, err = bu.Get()
	if len(bu.data) != 1 || bu.SizeNew() != 0 {
		t.Fatal("NEED TO BE BUFFER=1 AND FOLDER=1 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}

	fmt.Println("Adding Item")
	err = bu.Add(&entryExample)
	time.Sleep(100 * time.Millisecond)
	if len(bu.data) != 2 || bu.SizeNew() != 0 {
		t.Fatal("NEED TO BE BUFFER=1 AND FOLDER=1 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}

	fmt.Println("Adding Item")
	err = bu.Add(&entryExample)
	time.Sleep(100 * time.Millisecond)
	if len(bu.data) != 2 || bu.SizeNew() != 1 {
		t.Fatal("NEED TO BE BUFFER=1 AND FOLDER=1 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}

	fmt.Println("Reading Item")
	// Buffer will be full and re-read will not add new entries
	bu.ScanFolder()
	time.Sleep(100 * time.Millisecond)
	if len(bu.data) != 2 || bu.SizeNew() != 1 {
		t.Fatal("NEED TO BE BUFFER=2 AND FOLDER=2 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}

	fmt.Println("Adding Item")
	err = bu.Add(&entryExample)
	time.Sleep(100 * time.Millisecond)
	if len(bu.data) != 2 || bu.SizeNew() != 2 {
		t.Fatal("NEED TO BE BUFFER=2 AND FOLDER=5 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}

	fmt.Println("Adding Item")
	err = bu.Add(&entryExample)
	time.Sleep(100 * time.Millisecond)
	if len(bu.data) != 2 || bu.SizeNew() != 3 {
		t.Fatal("NEED TO BE BUFFER=2 AND FOLDER=5 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}

	fmt.Println("Adding Item")
	err = bu.Add(&entryExample)
	time.Sleep(100 * time.Millisecond)
	if len(bu.data) != 2 || bu.SizeNew() != 4 {
		t.Fatal("NEED TO BE BUFFER=2 AND FOLDER=5 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}

	fmt.Println("Getting Item")
	_, err = bu.Get()
	time.Sleep(100 * time.Millisecond)
	if len(bu.data) != 1 || bu.SizeNew() != 4 {
		t.Fatal("NEED TO BE BUFFER=1 AND FOLDER=4 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}

	fmt.Println("Getting Item")
	_, err = bu.Get()
	time.Sleep(100 * time.Millisecond)
	if len(bu.data) != 0 || bu.SizeNew() != 4 {
		t.Fatal("NEED TO BE BUFFER=1 AND FOLDER=3 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}

	fmt.Println("Getting Item")
	_, err = bu.Get()
	time.Sleep(200 * time.Millisecond)
	if len(bu.data) != 2 || bu.SizeNew() != 2 {
		t.Fatal("NEED TO BE BUFFER=2 AND FOLDER=2 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}
	fmt.Println("Getting Item")
	_, err = bu.Get()
	time.Sleep(100 * time.Millisecond)
	if len(bu.data) != 0 || bu.SizeNew() != 3 || !errors.Is(err, errorBufferIsEmpty) {
		t.Fatal("NEED TO BE BUFFER=0 AND FOLDER=5 , BUFFER = ", len(bu.data), " RETURN EMPTY => FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}

	fmt.Println("Reading Item")
	bu.ScanFolder()
	time.Sleep(300 * time.Millisecond)
	// Adding 3 to folder  Getting One buffer and Scan Folder
	if len(bu.data) != 2 || bu.SizeNew() != 1 {
		t.Fatal("NEED TO BE BUFFER=2 AND FOLDER=3 , BUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	} else {
		fmt.Println("\tBUFFER = ", len(bu.data), "FOLDER=", bu.SizeNew(), "ERROR:", err)
	}
}
