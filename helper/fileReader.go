package helper

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type FileStruct struct {
	InputFileName  string
	OutputFileName string
}

func (f *FileStruct) ReadFile() (string, error) {
	file, err := os.Open(f.InputFileName)
	if err != nil {
		return "", fmt.Errorf("Could not open file %s, because %v", f.InputFileName, err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	var store strings.Builder
	for {
		n, err := reader.Read(buffer)
		if n > 0 {
			store.WriteString(string(buffer[:n]))
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return "", fmt.Errorf("Could not Read file %s, because %v", f.InputFileName, err)
		}
	}
	return store.String(), nil
}

func (f *FileStruct) WriteTo(text string) error {
	file, err := os.Create(f.OutputFileName)

	if err != nil {
		return fmt.Errorf("Could not Create file %s, because %v", f.OutputFileName, err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	content := text
	_, err = writer.WriteString(content)
	if err != nil {
		return fmt.Errorf("Could not Write file %s, because %v", f.OutputFileName, err)
	}

	writer.Flush()
	return nil
}
