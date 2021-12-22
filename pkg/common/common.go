package common

import (
	"os"
	"path/filepath"
)

// 判断所给路径文件/文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// 获取文件夹下的文件夹
func GetDirs(path string) ([]string, error) {
	if !IsExist(path) {
		return []string{}, nil
	}
	var dirs []string
	err := filepath.Walk(path,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				dirs = append(dirs, path)
				return nil
			}

			return nil
		})
	return dirs, err
}
