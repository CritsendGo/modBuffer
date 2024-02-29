package modBuffer

import "errors"

// Error
var (
	errorFolderUnset      = errors.New("unset folder BufferFolder , please set it before init the buffer")
	errorBufferIsFull     = errors.New("buffer is full , stock it in BufferFolder")
	errorBufferIsEmpty    = errors.New("buffer is empty ")
	errorBufferAlreadySet = errors.New("buffer is empty ")
)

// Public Vars
var (
	BufferFolder string
	Debug        bool
)

// Private Vars
var (
	bufferList map[string]bool
)
