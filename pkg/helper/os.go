package helper

import (
	"errors"
	"os"
)

func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, errors.New("文件已存在")
	} else {
		return false, err
	}
}
