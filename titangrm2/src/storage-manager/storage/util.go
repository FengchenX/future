package storage

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"

	"grm-service/crypto"

	"github.com/globalsign/mgo"
	_ "github.com/jackc/pgx/stdlib"

	"grm-service/dbcentral/pg"

	"storage-manager/types"
)

type DeviceInfo struct {
	MountPath   string
	FileSystem  string
	Volume      string
	Used        string
	Avail       string
	UsedPercent string
}

// 获取挂在目录所在分区的情况
func GetDeviceInfo(mountPath string) (*DeviceInfo, error) {
	command := "df"
	params := []string{mountPath, "-hP"}
	cmd := exec.Command(command, params...)
	// 显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	// 实时循环读取输出流中的一行内容
	reader := bufio.NewReader(stdout)
	reader.ReadString('\n')

	line, err := reader.ReadString('\n')
	if err != nil || io.EOF == err {
		return nil, types.ErrGetDeviceInfo
	}
	fmt.Println("Line:", line)
	data := strings.Fields(line)
	if len(data) < 5 {
		return nil, types.ErrGetDeviceInfo
	}

	volume := strings.Replace(data[1], " ", "", -1)
	volume = strings.ToUpper(volume)
	if strings.HasSuffix(volume, "K") ||
		strings.HasSuffix(volume, "M") ||
		strings.HasSuffix(volume, "G") ||
		strings.HasSuffix(volume, "T") {
		volume = volume + "B"
	}

	used := strings.Replace(data[2], " ", "", -1)
	used = strings.ToUpper(used)
	if strings.HasSuffix(used, "K") ||
		strings.HasSuffix(used, "M") ||
		strings.HasSuffix(used, "G") ||
		strings.HasSuffix(used, "T") {
		used = used + "B"
	}

	avail := strings.Replace(data[3], " ", "", -1)
	avail = strings.ToUpper(avail)
	if strings.HasSuffix(avail, "K") ||
		strings.HasSuffix(avail, "M") ||
		strings.HasSuffix(avail, "G") ||
		strings.HasSuffix(avail, "T") {
		avail = avail + "B"
	}

	info := &DeviceInfo{
		FileSystem:  data[0],
		Volume:      volume,
		Used:        used,
		Avail:       avail,
		UsedPercent: data[4],
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	return info, nil
}

// 获取TitanCloud.Data数据库使用情况
func GetDeviceDBInfo(dev *types.Device) (*DeviceInfo, error) {
	pwd, err := crypto.AesDecrypt(dev.DBPwd)
	if err != nil {
		return nil, err
	}

	dataSource := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dev.DBUser, pwd, dev.IpAddress, dev.DBPort, pg.DataDBName)
	fmt.Println(dataSource)
	db, err := sql.Open("pgx", dataSource)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var info *DeviceInfo
	sql := fmt.Sprintf(`select pg_database_size('%s'),pg_size_pretty(pg_database_size('%s'))`,
		pg.DataDBName, pg.DataDBName)
	rows, err := db.Query(sql)
	if err != nil {
		return info, err
	}
	defer rows.Close()

	var Used int64
	var UsedP string
	if rows.Next() {
		err = rows.Scan(&Used, &UsedP)
		if err != nil {
			return info, err
		}
	}
	if err := rows.Err(); err != nil {
		return info, err
	}

	volume := strings.Replace(dev.Volume, " ", "", -1)
	volume = strings.ToUpper(volume)
	types := map[string]uint32{
		"KB": 10,
		"MB": 20,
		"GB": 30,
		"TB": 40,
		"K":  10,
		"M":  20,
		"G":  30,
		"T":  40,
	}

	var volInt int64
	for k, v := range types {
		if index := strings.LastIndex(volume, k); index != -1 {
			vol, err := strconv.ParseInt(volume[:index], 10, 64)
			if err != nil {
				return info, err
			}
			volInt = vol << v
			break
		}
	}

	avail := volInt - Used
	var quote string
	if avail > (1 << 40) {
		avail = avail / (1 << 40)
		quote = "TB"
	} else if avail > (1 << 30) {
		avail = avail / (1 << 30)
		quote = "GB"
	} else if avail > (1 << 20) {
		avail = avail / (1 << 20)
		quote = "MB"
	} else if avail > (1 << 10) {
		avail = avail / (1 << 10)
		quote = "KB"
	}

	fmt.Println("total:", volInt, ",used:", Used)
	UsedP = strings.Replace(UsedP, " ", "", -1)

	usedPercent := float64(Used) / float64(volInt) * 100
	info = &DeviceInfo{
		Volume:      volume,
		Used:        UsedP,
		Avail:       strconv.FormatInt(avail, 10) + quote,
		UsedPercent: strconv.FormatFloat(usedPercent, 'f', 0, 64) + "%",
	}
	return info, nil
}

// 获取mongodb数据库使用情况
type dbNames struct {
	Databases []struct {
		Name       string `bson:"name"`
		Empty      bool   `bson:"empty"`
		SizeOnDisk int64  `bson:"sizeOnDisk"`
	} `bson:"databases"`
}

func GetDeviceMongoInfo(dev *types.Device) (*DeviceInfo, error) {
	url := fmt.Sprintf("mongodb://%s:%s", dev.IpAddress, dev.DBPort)
	session, err := mgo.Dial(url)
	if err != nil {
		fmt.Println("Dial: ", err)
		return nil, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	var result dbNames
	if err = session.Run("listDatabases", &result); err != nil {
		fmt.Println("listDatabases: ", err)
		return nil, err
	}

	var Used int64
	var UsedP string
	for _, val := range result.Databases {
		if !val.Empty {
			Used += val.SizeOnDisk
		}
	}

	volume := strings.Replace(dev.Volume, " ", "", -1)
	volume = strings.ToUpper(volume)
	types := map[string]uint32{
		"KB": 10,
		"MB": 20,
		"GB": 30,
		"TB": 40,
		"K":  10,
		"M":  20,
		"G":  30,
		"T":  40,
	}

	var volInt int64
	for k, v := range types {
		if index := strings.LastIndex(volume, k); index != -1 {
			vol, err := strconv.ParseInt(volume[:index], 10, 64)
			if err != nil {
				return nil, err
			}
			volInt = vol << v
			break
		}
	}

	avail := volInt - Used
	var quote string
	if avail > (1 << 40) {
		avail = avail / (1 << 40)
		quote = "TB"
	} else if avail > (1 << 30) {
		avail = avail / (1 << 30)
		quote = "GB"
	} else if avail > (1 << 20) {
		avail = avail / (1 << 20)
		quote = "MB"
	} else if avail > (1 << 10) {
		avail = avail / (1 << 10)
		quote = "KB"
	}

	fmt.Println("total:", volInt, ",used:", Used)
	UsedP = strings.Replace(UsedP, " ", "", -1)
	usedPercent := float64(Used) / float64(volInt) * 100

	//
	var Usedquote string
	if Used > (1 << 40) {
		Used = Used / (1 << 40)
		Usedquote = "TB"
	} else if Used > (1 << 30) {
		Used = Used / (1 << 30)
		Usedquote = "GB"
	} else if Used > (1 << 20) {
		Used = Used / (1 << 20)
		Usedquote = "MB"
	} else if Used > (1 << 10) {
		Used = Used / (1 << 10)
		Usedquote = "KB"
	}
	info := &DeviceInfo{
		Volume:      volume,
		Used:        strconv.FormatInt(Used, 10) + Usedquote,
		Avail:       strconv.FormatInt(avail, 10) + quote,
		UsedPercent: strconv.FormatFloat(usedPercent, 'f', 0, 64) + "%",
	}
	return info, nil
}
