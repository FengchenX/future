package compress

import (
	"archive/tar"
	"compress/gzip"
	"grm-service/util"
	"io"
	"os"
	"path/filepath"
)

func Untar(file, dest string) error {
	fr, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fr.Close()
	// gzip read
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer gr.Close()
	// tar read
	tr := tar.NewReader(gr)
	// 读取文件
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		untarPath := filepath.Join(dest, h.Name)
		if h.FileInfo().IsDir() {
			util.CheckDir(untarPath)
			continue
		}
		util.CheckDir(filepath.Dir(untarPath))
		fw, err := os.Create(untarPath)
		if err != nil {
			return err
		}
		defer fw.Close()
		// 写文件
		_, err = io.Copy(fw, tr)
		if err != nil {
			return err
		}
		fw.Close()
	}
	return nil
}
