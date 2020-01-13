package utils

import (
	"errors"
	"os"
)

func WriteFile(filename, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return errors.New("file not created")
	}
	file.WriteString(data)
	file.Close()
	return nil
}
