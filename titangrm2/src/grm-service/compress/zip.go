package compress

import (
	"archive/zip"
	"fmt"
	"grm-service/util"
	"io"
	"os"
	"path/filepath"
)

func Unzip(file, dest string) error {
	fmt.Printf("[INFO]start unzip %s to %s\n", file, dest)
	File, err := zip.OpenReader(file)
	if err != nil {
		return err
	}
	defer File.Close()
	for _, v := range File.File {
		//fmt.Println("before: ", v.Name)
		v.Name = ConvertToString(v.Name, "gbk", "utf-8")
		//fmt.Println("after: ", v.Name)
		if v.FileInfo().IsDir() {
			_path := filepath.Join(dest, v.Name)
			util.CheckDir(_path)
			continue
		}

		srcFile, err := v.Open()
		if err != nil {
			return err
		}
		defer srcFile.Close()

		unzipPath := filepath.Join(dest, v.Name)
		util.CheckDir(filepath.Dir(unzipPath))
		fw, err := os.OpenFile(unzipPath, os.O_CREATE|os.O_WRONLY, os.ModePerm /*os.FileMode(h.Mode)*/)
		if err != nil {
			return err
		}
		defer fw.Close()

		// 写文件
		_, err = io.Copy(fw, srcFile)
		if err != nil {
			return err
		}
	}
	fmt.Println("[INFO]unzip finish")
	return nil
}
