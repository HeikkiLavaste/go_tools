package gotools

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// randomStringSource is the source for generating random strings.
const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321_+"

// defaultMaxUpload is the default max upload size (10 mb)
const defaultMaxUpload = 10485760

type Tools struct {
	MaxJSONSize        int         // maximum size of JSON file we'll process
	MaxXMLSize         int         // maximum size of XML file we'll process
	MaxFileSize        int         // maximum size of uploaded files in bytes
	AllowedFileTypes   []string    // allowed file types for upload (e.g. image/jpeg)
	AllowUnknownFields bool        // if set to true, allow unknown fields in JSON
	ErrorLog           *log.Logger // the info log.
	InfoLog            *log.Logger // the error log.
}

func New() Tools {
	return Tools{
		MaxJSONSize: defaultMaxUpload,
		MaxXMLSize:  defaultMaxUpload,
		MaxFileSize: defaultMaxUpload,
		InfoLog:     log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog:    log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (t *Tools) ReadCSV(src string) ([]byte, error) {
	if !IsOfType(src, "csv") {
		t.ErrorLog.Println("Not CSV")
		return nil, nil
	}
	data, err := os.ReadFile(src)
	if err != nil {
		t.ErrorLog.Println(err.Error())
	}
	fmt.Println(IsOfType(src, "csv"))
	return data, nil

}

func IsOfType(name string, ftype string) bool {
	return filepath.Ext(name) == "."+ftype
}
