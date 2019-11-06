package utils

import (
	"log"
	"os"
	"path/filepath"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

var fd *os.File
var err error

func OpenFile(fullPath string) *os.File {
	if PathExists(fullPath) {
		fd, err = os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY, 0777)
		if err != nil {
			log.Fatal("createFile error1:", err)
		}
	} else {
		dir:=filepath.Dir(fullPath)
		_, err := os.Stat(dir)
		if err != nil {
			pathAll,_:=filepath.Abs(dir)
			err1:=os.MkdirAll(pathAll,0777)
			if err1 != nil {
				log.Fatal("createFile error3:", err1)
			}
		}
		fd, err = os.Create(fullPath)
		if err != nil {
			log.Fatal("createFile error4:", err)
		}
	}
	return fd
}
