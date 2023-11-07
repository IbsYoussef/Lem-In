package methods

import (
	"os"
)

func OpenFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return "", err
	}
	fileSize := stat.Size()

	fileContent := make([]byte, fileSize)

	_, err = file.Read(fileContent)
	if err != nil {
		return "", err
	}

	fileText := string(fileContent)

	return fileText, nil
}
