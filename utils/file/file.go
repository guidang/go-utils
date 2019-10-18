package file

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// PathExists 文件(夹)路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetCurrentDirectory 获取当前路径
func GetCurrentDirectory() (path string, err error) {
	path, err = filepath.Abs(filepath.Dir(os.Args[0]))
	return
}

// MD5 文件 MD5
func MD5(file *multipart.FileHeader) (md5Str string, err error) {
	f, err := file.Open()
	if err != nil {
		return
	}
	defer f.Close()

	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
