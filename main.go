package modBuffer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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
	err := b.Init()
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
	if os.MkdirAll(b.folder+"new/", os.ModePerm) != nil {
		fmt.Println("Error creating folder:", b.folder+"new/")
	}
	/// err for CSBuffer.Error()
	if os.MkdirAll(b.folder+"err/", os.ModePerm) != nil {
		fmt.Println("Error creating folder:", b.folder+"err/")
	}
	/// ack for CSBuffer.Finish()
	if os.MkdirAll(b.folder+"ack/", os.ModePerm) != nil {
		fmt.Println("Error creating folder:", b.folder+"ack/")
	}
	return nil
}

func (b *CSBuffer) Add(e *any) error {
	if Debug == true {
		fmt.Printf("%+v\n", b)
	}
	b.mutex.Lock()
	defer b.mutex.Unlock()
	// Check if buffer is full
	if len(b.data) >= b.maxSize {
		// Write on disk event return on success critical on error
		err := b.Save(e)
		if err != nil {
			log.Println("BUFFER FULL AND UNABLE TO SAVE EVENT ON DISK")
			return err
		}
		return errorBufferIsFull
	}
	b.data = append(b.data, e)
	return nil
}
func (b *CSBuffer) Get() (any, error) {
	// Lock Usage
	b.mutex.Lock()
	defer b.mutex.Unlock()
	// Check if buffer is empty
	if len(b.data) == 0 {
		return nil, errorBufferIsEmpty
	}
	// Get and remove the first item from the buffer
	item := b.data[0]
	b.data = b.data[1:]
	return item, nil
}
func (b *CSBuffer) Save(e *any) error {
	if Debug == true {
		fmt.Printf("%+v\n", b)
	}
	fileName := fmt.Sprintln(time.Now().UnixMicro())
	filePath := BufferFolder + fileName
	f, err := os.Create(filePath)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println("ERROR IN CLOSING FILE", err)
		}
	}(f)
	if err != nil {
		log.Println("CREATE EVENT DD", err)
		return err
	}
	eBit, err := json.Marshal(e)
	if err != nil {
		log.Println("JSON EVENT DD", err)
		return err
	}
	_, err = f.Write(eBit)
	if err != nil {
		log.Println("WRITE EVENT DD", err)
		return err
	}
	return nil
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
