package importer

import (
	"encoding/json"
	"fmt"
	"grm-service/compress"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	grmpath "grm-service/path"
)

// 获取系统FileSysDomain配置
func getFileSysDomain(configFile string) (map[string]string, error) {
	var domains map[string]string
	domains = make(map[string]string)

	var config []struct {
		Name string `json:"name"`
		Path string `json:"path"`
	}
	file, err := os.Open(configFile)
	if err != nil {
		return domains, fmt.Errorf("%s: %s", err.Error(), configFile)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return domains, err
	}
	for _, value := range config {
		domains[value.Name] = value.Path
	}
	return domains, nil
}

// 遍历
func walkDir(dirPth, domainKey, domainPath string) (sysDomains, error) {
	var ret sysDomains

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	for _, fi := range dir {
		path := filepath.Join(dirPth, fi.Name())
		domain := domain{
			Name:  fi.Name(),
			Path:  path,
			IsDir: fi.IsDir(),
		}
		if strings.HasPrefix(path, domainPath) {
			domain.Domain = strings.Replace(domain.Path, domainPath, domainKey, -1)
		}
		ret = append(ret, domain)
	}
	return ret, err
}

func writeLoadFiles(file string, fileList []string) error {
	fileP, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fileP.Close()

	for _, id := range fileList {
		if !strings.HasSuffix(id, "\n") {
			id = id + "\n"
		}
		if _, err := io.WriteString(fileP, id); err != nil {
			return err
		}
	}
	return nil
}

// 上传文件并解压
func uploadFormFile(dataFile string, file io.Reader) (string, error) {
	// 写文件
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	filePath, fileName := filepath.Split(dataFile)
	if err := grmpath.CreateAllPath(filePath); err != nil {
		return "", err
	}
	fmt.Println("upload dataFile:", dataFile)
	//fmt.Println("upload dir:", filePath)
	//fmt.Println("upload file:", fileName)
	if err := ioutil.WriteFile(dataFile, content, os.ModePerm); err != nil {
		return "", err
	}

	// 解压
	dataName := strings.ToLower(fileName)
	dir := filepath.Join(filePath, dataName[:strings.Index(dataName, ".")])
	if strings.HasSuffix(dataName, ".zip") {
		grmpath.CreateAllPath(dir)
		filePath = dir
		if err := compress.Unzip(dataFile, dir); err != nil {
			return "", err
		}
	} else if strings.HasSuffix(dataName, ".tar.gz") {
		grmpath.CreateAllPath(dir)
		filePath = dir
		if err := compress.Untar(dataFile, dir); err != nil {
			return "", err
		}
	}
	return filePath, nil
}
