package utils

import (
	"bufio"
	"errors"
	"io"
	"os"
)

// ReadFile 读取文件
func ReadFile(filepath string) ([]string, error) {
	if filepath == "" {
		return nil, errors.New("Utils.ReadFile filepath is empty")
	}
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := []string{}

	line := bufio.NewReader(file)
	for {
		content, _, err := line.ReadLine()
		if err == io.EOF {
			return result, nil
		}
		result = append(result, string(content))
	}
}
