package modBuffer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func init() {
	bufferList = make(map[string]bool)
}

// NewBuffer Used To Create New Buffer and Init it
func NewBuffer(folder string, size int) (*CSBuffer, error) {
	if bufferList[folder] == true {
		return nil, errorBufferAlreadySet
	}
	b := CSBuffer{
		maxSize: size,
		mutex:   sync.Mutex{},
		folder:  folder,
	}
	if Debug == true {
		fmt.Println("DEBUG: NEW BUFFER FOR ", folder, "WITH SIZE", size)
	}
	err := b.Init()
	// Read Folder Each Interval to refill buffer
	go b.Fill()

	// No Error lock buffer Set
	if err != nil {
		bufferList[folder] = true
	}
	return &b, err
}

func (b *CSBuffer) Init() error {
	if b.folder == "" {
		return errorFolderUnset
	}
	// Create Sub Folder
	/// new for CSBuffer.Save() and CSBuffer.Read()function
	if os.MkdirAll(b.folder+"new/", 777) != nil {
		fmt.Println("Error creating folder:", b.folder+"new/")
	}
	/// err for CSBuffer.Error()
	if os.MkdirAll(b.folder+"err/", 777) != nil {
		fmt.Println("Error creating folder:", b.folder+"err/")
	}
	/// ack for CSBuffer.Finish()
	if os.MkdirAll(b.folder+"ack/", 777) != nil {
		fmt.Println("Error creating folder:", b.folder+"ack/")
	}
	go b.Survey()
	return nil
}

func (b *CSBuffer) Add(e any) error {
	if Debug == true {
		fmt.Println("DEBUG:", "ADD NEW SIZE => ", len(b.data))
	}
	var item *any = &e
	b.mutex.Lock()
	defer b.mutex.Unlock()
	// Check if buffer is full
	if len(b.data) >= b.maxSize {
		if Debug == true {
			fmt.Println("DEBUG:", "SIZE IS OVER", b.maxSize)
		}
		// Write on disk event return on success critical on error

		err := b.Save(item)
		if err != nil {
			log.Println("ERROR:BUFFER FULL AND UNABLE TO SAVE EVENT ON DISK")
			return err
		}
		return errorBufferIsFull
	}
	b.data = append(b.data, item)
	return nil
}

func (b *CSBuffer) Fill() {
	for {
		if Debug == true {

		}
		if len(b.data) != b.maxSize && b.SizeNew() != 0 {
			b.ScanFolder()
		}
		time.Sleep(500 * time.Millisecond)
	}
}
func (b *CSBuffer) Survey() {
	for {
		if Debug == true {
			fmt.Println("SURVEY SIZE OF POOL ", len(b.data), " SIZE OF FOLDER ", b.SizeNew(), b)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
func (b *CSBuffer) Get() (any, error) {
	// Lock Usage
	b.mutex.Lock()
	// Check if buffer is empty
	if len(b.data) == 0 {
		b.mutex.Unlock()
		return nil, errorBufferIsEmpty
	}
	// Get and remove the first item from the buffer
	item := b.data[0]
	b.data = b.data[1:]
	b.mutex.Unlock()
	return item, nil
}
func (b *CSBuffer) Save(e *any) error {

	if Debug == true {
		fmt.Println("DEBUG:", "SAVING ITEM TO FOLDER", b.folder+"new/")
	}
	fileName := fmt.Sprintln(time.Now().UnixMicro())
	fileName = strings.TrimSpace(fileName)
	data, err := json.Marshal(e)
	if err != nil {
		log.Println("ERROR MARSHAL:", err)
	}
	filePathItem := b.folder + "new/" + fileName + ".json"
	if Debug == true {
		fmt.Println("DEBUG:", filePathItem)
	}
	err = ioutil.WriteFile(filePathItem, data, 0660)
	if err != nil {
		log.Println("ERROR WRITE FILE:", err)
	}
	return nil
}
func (b *CSBuffer) SizeNew() int {
	var nb = 0
	if Debug == true {
		fmt.Println("DEBUG:", "CHECKING SIZE OF FOLDER NEW")
	}
	_ = filepath.WalkDir(b.folder+"new/", func(filePath string, d fs.DirEntry, err error) error {
		if strings.Index(d.Name(), ".json") > -1 {
			nb++
		}
		return nil
	})
	if Debug == true {
		//fmt.Println("DEBUG:", "SIZE IN NEW", nb)
	}
	return nb
}
func (b *CSBuffer) Read(filePath string) error {
	if Debug == true {
		fmt.Println("READING FILE", filePath)
	}

	return nil
}
func (b *CSBuffer) Error(content any) error {
	if Debug == true {
		fmt.Println("ERROR BUFFER", content)
	}

	return nil
}
func (b *CSBuffer) Finish(content any) error {
	if Debug == true {
		fmt.Println("FINISH BUFFER", content)
	}

	return nil
}

func (b *CSBuffer) ScanFolder() {
	err := filepath.WalkDir(b.folder+"new/", func(filePath string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		} else {
			if Debug {
				fmt.Println("DEBUG: READING ", b.folder+"new/"+d.Name())
			}
			filePathDetail := b.folder + "new/" + d.Name()
			bt, err := os.ReadFile(filePathDetail)
			if err != nil {
				return err
			}
			var obj any
			err = json.Unmarshal(bt, &obj)
			if err != nil {
				if Debug {
					fmt.Println("DEBUG: ERROR  ", err)
				}
				return err
			}
			err = b.Add(&obj)
			if err == nil {
				_ = os.Rename(filePathDetail, b.folder+"ack/"+d.Name())
			}
			if Debug {
				fmt.Println("DEBUG: ERROR  ", err)
				fmt.Println("DEBUG: SIZE  ", len(b.data))
			}
			return err
		}
	})
	if err != nil {
		if errors.Is(err, errorBufferIsFull) {
			return
		} else {

		}

	}

}
