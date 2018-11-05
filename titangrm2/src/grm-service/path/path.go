package path

import (
	"os"
)

// 判断是否为路径
func IsDir(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	return file.IsDir()
}

// 判断文件或路径是否存在
func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 创建路径
func CreateAllPath(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
