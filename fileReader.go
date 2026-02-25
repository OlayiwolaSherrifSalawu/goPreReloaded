package main

import (
	"bufio"
	"fmt"
	"os"
)

type FileStruct struct {
	FileName string
	errors   error
}

func (f *FileStruct) ReadFile() error {
	file, err := os.OpenFile(f.FileName, os.O_CREATE, 0644)
	if err != nil {
		f.errors = err
		return fmt.Errorf("Could not open file %s, because %v", f.FileName, f.errors)

	}
	defer file.Close()
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Print(string(buffer[:n]))
	}
	return nil
}
