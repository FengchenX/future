package util

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"crypto/md5"
	crand "crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"grm-service/common"
)

// 获取系统FileSysDomain配置
func GetFileSysDomain(configFile string) (map[string]string, error) {
	var domains map[string]string
	domains = make(map[string]string)

	var config []struct {
		Name string `json:"name"`
		Path string `json:"path"`
	}
	file, err := os.Open(configFile)
	if err != nil {
		return domains, err
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

func GetFileName(name, scp string) (string, bool) {
	//获取环境变量
	shares := strings.Split(scp, ";")
	for _, share := range shares {
		s := strings.Split(share, "=")

		if strings.HasPrefix(name, "$["+s[0]+"]") {
			name = strings.Replace(name, "$["+s[0]+"]", s[1], -1)
			return name, true
		}
	}
	return name, false
}

func IgnoreUTF8BOM(data []byte) []byte {
	if data == nil {
		return nil
	}
	if len(data) >= 3 && data[0] == 0xef && data[1] == 0xbb && data[2] == 0xbf {
		return data[3:]
	}
	return data
}

func GetGpuCount() int {
	var smiinfo string
	if runtime.GOOS == "windows" {
		cmd := exec.Command("C:\\Program Files\\NVIDIA Corporation\\NVSMI\\nvidia-smi.exe", "-a")
		bytes, err := cmd.Output()
		if err != nil {
			fmt.Println("GetGpuCount error:", err)
			return 0
		}

		smiinfo = string(bytes)
	} else if runtime.GOOS == "linux" {
		cmd := exec.Command("nvidia-smi", "-a")
		bytes, err := cmd.Output()
		if err != nil {
			fmt.Println("GetGpuCount error:", err)
			return 0
		}
		smiinfo = string(bytes)
	}
	infos := strings.Split(smiinfo, "\n")
	for _, info := range infos {
		if strings.HasPrefix(info, "Attached GPUs") {
			values := strings.Split(info, ":")
			if len(values) >= 2 {
				count, err := strconv.Atoi(strings.TrimSpace(values[1]))
				if err != nil {
					return 0
				}
				return count
			}
		}
	}
	return 0
}

func MatchString(pattern, str string) (matched bool, err error) {
	//将pattern中的*,?转化为正则表达式内容，然后用正则表达式来处理
	pattern = strings.Replace(pattern, "*", ".*", -1)
	pattern = strings.Replace(pattern, "?", ".", -1)

	//	return "^" + Regex.Escape(pattern).Replace("\\*", ".*").Replace("\\?", ".") + "$"
	return regexp.MatchString(pattern, str)
}

func GetTempDir() string {
	dir := filepath.Join(os.TempDir(), common.Namespace, GenerateUUID())
	CheckDir(dir)
	return dir
}

func GetTaskidIndex(taskid string) string {
	s := strings.Split(taskid, ".")
	if len(s) == 2 {
		return s[1]
	} else {
		log.Println("[error]taskid is wrong:", taskid)
		return taskid
	}
}

//======================================================解压
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
		// 打开文件
		untarPath := filepath.Join(dest, h.Name)
		err = CheckFileDir(untarPath)
		if err != nil {
			return err
		}
		fw, err := os.OpenFile(untarPath, os.O_CREATE|os.O_WRONLY, os.ModePerm /*os.FileMode(h.Mode)*/)
		if err != nil {
			return err
		}
		defer fw.Close()
		// 写文件
		_, err = io.Copy(fw, tr)
		if err != nil {
			return err
		}
	}
	return nil
}

func Unzip(file, dest string) error {
	fmt.Printf("[INFO]start unzip %s to %s\n", file, dest)
	File, err := zip.OpenReader(file)
	if err != nil {
		return err
	}
	defer File.Close()
	for _, v := range File.File {
		if v.FileInfo().IsDir() {
			//目录
			_path := filepath.Join(dest, v.Name)
			err = CheckFileDir(_path)
			if err != nil {
				return err
			}
		} else {
			srcFile, err := v.Open()
			if err != nil {
				return err
			}
			defer srcFile.Close()
			unzipPath := filepath.Join(dest, v.Name)
			err = CheckFileDir(unzipPath)
			if err != nil {
				return err
			}
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
	}
	fmt.Println("[INFO]unzip finish")
	return nil
}

//=======================================================压缩
func ZipDir(srcDirPath string, recPath string, zw *zip.Writer) error {
	// Open source diretory
	dir, err := os.Open(srcDirPath)
	if err != nil {
		return err
	}
	defer dir.Close()

	// Get file info slice
	fis, err := dir.Readdir(0)
	if err != nil {
		return err
	}
	for _, fi := range fis {
		// Append path
		curPath := srcDirPath + "/" + fi.Name()
		// Check it is directory or file
		if fi.IsDir() {
			// Directory
			// (Directory won't add unitl all subfiles are added)
			//			fmt.Printf("Adding path...%s\n", curPath)
			ZipDir(curPath, recPath+"/"+fi.Name(), zw)
		} else {
			// File
			//			fmt.Printf("Adding file...%s\n", curPath)
		}

		ZipFile(curPath, recPath+"/"+fi.Name(), zw, fi)
	}
	return nil
}

func ZipFile(srcFile string, recPath string, zw *zip.Writer, fi os.FileInfo) error {
	if fi.IsDir() {
		// Create tar header
		//hdr := new(zip.FileHeader)
		hdr := &zip.FileHeader{
			Name:   recPath + "/",
			Flags:  1 << 11, // 使用utf8编码
			Method: zip.Deflate,
		}

		// if last character of header name is '/' it also can be directory
		// but if you don't set Typeflag, error will occur when you untargz
		//hdr.Name = recPath + "/"
		//		hdr.Typeflag = tar.TypeDir
		//		hdr.Size = 0
		//hdr.Mode = 0755 | c_ISDIR
		//		hdr.Mode = int64(fi.Mode())
		//		hdr.ModTime = fi.ModTime()

		// Write hander
		_, err := zw.CreateHeader(hdr)
		if err != nil {
			return err
		}
	} else {
		// File reader
		fr, err := os.Open(srcFile)
		if err != nil {
			return err
		}
		defer fr.Close()

		// Create tar header
		//		hdr := new(zip.FileHeader)
		//		hdr.Name = recPath
		//		hdr.Size = fi.Size()
		//		hdr.Mode = int64(fi.Mode())
		//		hdr.ModTime = fi.ModTime()

		hdr := &zip.FileHeader{
			Name:   recPath,
			Flags:  1 << 11, // 使用utf8编码
			Method: zip.Deflate,
		}

		// Write hander
		w, err := zw.CreateHeader(hdr)
		if err != nil {
			return err
		}

		// Write file data
		_, err = io.Copy(w, fr)
		if err != nil {
			return err
		}
	}
	return nil
}

func ZipFiles(frm, dst string) error {
	buf := bytes.NewBuffer(make([]byte, 0, 10*1024*1024)) // 创建一个读写缓冲
	myzip := zip.NewWriter(buf)                           // 用压缩器包装该缓冲
	// 用Walk方法来将所有目录下的文件写入zip
	err := filepath.Walk(frm, func(path string, info os.FileInfo, err error) error {
		var file []byte
		if err != nil {
			return filepath.SkipDir
		}
		header, err := zip.FileInfoHeader(info) // 转换为zip格式的文件信息
		if err != nil {
			return filepath.SkipDir
		}
		header.Name, _ = filepath.Rel(filepath.Dir(frm), path)
		if !info.IsDir() {
			// 确定采用的压缩算法（这个是内建注册的deflate）
			header.Method = 8
			file, err = ioutil.ReadFile(path) // 获取文件内容
			if err != nil {
				return filepath.SkipDir
			}
		} else {
			file = nil
		}
		// 上面的部分如果出错都返回filepath.SkipDir
		// 下面的部分如果出错都直接返回该错误
		// 目的是尽可能的压缩目录下的文件，同时保证zip文件格式正确
		w, err := myzip.CreateHeader(header) // 创建一条记录并写入文件信息
		if err != nil {
			return err
		}
		_, err = w.Write(file) // 非目录文件会写入数据，目录不会写入数据
		if err != nil {        // 因为目录的内容可能会修改
			return err // 最关键的是我不知道咋获得目录文件的内容
		}
		return nil
	})
	if err != nil {
		return err
	}
	myzip.Close()               // 关闭压缩器，让压缩器缓冲中的数据写入buf
	file, err := os.Create(dst) // 建立zip文件
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = buf.WriteTo(file) // 将buf中的数据写入文件
	if err != nil {
		return err
	}
	return nil
}

//获取路径中的文件名，不包含后缀
func GetFileNameWithoutExt(path string) string {
	filenameWithSuffix := filepath.Base(path)
	fileSuffix := filepath.Ext(filenameWithSuffix)
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return filenameOnly
}

func GetFileNameWithExt(path string) string {
	filenameWithSuffix := filepath.Base(path)
	return filenameWithSuffix
}

// GenerateUUID is used to generate a random UUID
func GenerateUUID() string {
	buf := make([]byte, 16)
	if _, err := crand.Read(buf); err != nil {
		fmt.Errorf("failed to read random bytes: %v", err)
	}

	return fmt.Sprintf("%08x%04x%04x%04x%12x",
		buf[0:4],
		buf[4:6],
		buf[6:8],
		buf[8:10],
		buf[10:16])
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func CreateSessionId() string {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}

func Str2float(str string) float64 {
	if bar, err := strconv.ParseFloat(str, 10); err == nil {
		return bar
	} else {
		return common.MinInt
	}
}

func Str2int(str string) int {
	if bar, err := strconv.Atoi(str); err == nil {
		return bar
	} else {
		return common.MinInt
	}
}

func Str2int64(str string) int64 {
	if bar, err := strconv.ParseInt(str, 10, 64); err == nil {
		return bar
	} else {
		return common.MinInt
	}
}

//func Int2bool(i int) bool {
//	if i == 1 {
//		return true
//	} else {
//		return false
//	}
//}

func Str2BoolFalse(str string) bool {
	if bar, err := strconv.ParseBool(str); err != nil {
		return false
	} else {
		return bar
	}
}

func Str2BoolTrue(str string) bool {
	if bar, err := strconv.ParseBool(str); err != nil {
		return true
	} else {
		return bar
	}
}

func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && strings.TrimSpace(a[i-1]) == strings.TrimSpace(a[i])) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

func GetRandom(list []string) string {
	return list[rand.Intn(len(list))]
}

func Interrupt() error {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return fmt.Errorf("%s", <-c)
}

//time相关操作
func GetDayBefore(day int) string {
	now := time.Now()
	nowEnd := time.Date(now.Year(), now.Month(), now.Day(), 24, 0, 0, 0, time.Local)

	daydiff, _ := time.ParseDuration(fmt.Sprintf("-%dh", day*24))
	daystart := nowEnd.Add(daydiff)
	dayStr := daystart.Format("2006-01-02")
	return dayStr
}

func GetTimeAfter(day int) string {
	now := time.Now()
	daydiff, _ := time.ParseDuration(fmt.Sprintf("+%dh", day*24))
	daystart := now.Add(daydiff)
	dayStr := daystart.Format("2006-01-02 15:04:05")
	return dayStr
}

func DBTimeStd(dbtime string) string {
	if dbtime == "1000-01-01T00:00:00Z" {
		return ""
	}
	t, err := time.Parse("2006-01-02T15:04:05Z", dbtime)
	if err != nil {
		return dbtime
	}

	return t.Format("2006-01-02 15:04:05")
}

func GetTimeStd(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func GetDateStd(t time.Time) string {
	return t.Format("2006-01-02")
}

func GetTimeNow() string {
	return time.Now().Format("20060102-150405")
}

func GetTimeNowStd() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func PgStr2Date(s string) string {
	return "to_date('" + s + "','YYYY-MM-DD')"
}

func PgStr2Time(s string) string {
	return "to_timestamp('" + s + "','YYYY-MM-DD hh24:mi:ss')"
}

func GetTimeNowDB() string {
	return "current_timestamp"
}

func GetTimeDB(date string) string {
	return "datetime('" + date + "')"
}

func GetDayNow() string {
	return time.Now().Format("20060102")
}

func PatchDay2Time(day string) time.Time {
	t, err := time.Parse("20060102", day)
	if err != nil {
		return time.Now()
	}
	return t
}

func String2Time(str string) time.Time {
	t, err := time.Parse("2006-01-02T15:04:05Z", str)
	if err != nil {
		return time.Now()
	}
	return t
}

func String2DayString(str string) string {
	t := String2Time(str)
	return t.Format("2006年01月02日")
}

func CopyFile(dstName, srcName string) (err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	return err
}

func CopyOrCreateFile(dstName, srcName string) (err error) {
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer dst.Close()

	src, errSrc := os.Open(srcName)
	if errSrc == nil {
		_, err = io.Copy(dst, src)
	}
	return err
}

//路径相关操作
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func CreateLogPath(key string) string {
	fileExe, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(fileExe)

	logDir := filepath.Join(filepath.Dir(path), "log")
	if !PathExists(logDir) {
		os.MkdirAll(logDir, os.ModePerm)
	}
	logPath := filepath.Join(logDir, key+"_"+GetTimeNow()+".log")
	return logPath
}

func CreateRaftDir(key string) string {
	fileExe, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(fileExe)

	raftDir := filepath.Join(filepath.Dir(path), "raft")
	if !PathExists(raftDir) {
		os.MkdirAll(raftDir, os.ModePerm)
	}
	dirPath := filepath.Join(raftDir, key)
	if !PathExists(dirPath) {
		os.MkdirAll(dirPath, os.ModePerm)
	}
	return dirPath
}

func CheckFileDir(fp string) error {
	dir := filepath.Dir(fp)
	if !PathExists(dir) {
		return os.MkdirAll(dir, os.ModePerm)
	}
	return nil
}

func WalkDir(dirPth string) (files []string, err error) {
	files = make([]string, 0)
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil { //忽略错误
			return err
		}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		files = append(files, filename)
		return nil
	})
	return files, err
}

func CheckDir(dp string) error {
	if !PathExists(dp) {
		return os.MkdirAll(dp, os.ModePerm)
	}
	return nil
}

func IsDir(file string) bool {
	f, e := os.Stat(file)
	if e != nil {
		return false
	}
	return f.IsDir()
}

func AllowOrigin(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers,auth-session,Accept, Content-Type, Content-Length, Accept-Encoding,X-CSRF-Token,Authorization,X-Requested-With")
}

func AllowSecurity(w http.ResponseWriter) {
	//	w.Header().Set("Content-Security-Policy", "")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
}

func Str2uint32(str string) uint32 {
	if bar, err := strconv.Atoi(str); err == nil {
		return uint32(bar)
	} else {
		return 0
	}
}
