package main

import (
	"os"
	"strconv"
)

// GenerateMessageID generates an incrementing message ID
func GenerateMessageID(fileName string) (string, error) {
	currentID := "0"
	file, err := os.Open(fileName)
	if err == nil {
		var buffer [128]byte
		n, err := file.Read(buffer[:])
		if err != nil {
			file.Close()
			return "", err
		}
		currentID = string(buffer[:n])
		file.Close()
	} else if !os.IsNotExist(err) {
		return "", err
	}

	nextIDInt, err := strconv.ParseInt(currentID, 10, 64)
	if err != nil {
		return "", err
	}

	nextID := strconv.FormatInt(nextIDInt + 1, 10)

	file, err = os.Create(fileName)
	if err != nil {
		return "", err
	}
	_, err = file.WriteString(nextID)
	if err != nil {
		file.Close()
		return "", err
	}

	file.Close()

	return currentID, nil
}
