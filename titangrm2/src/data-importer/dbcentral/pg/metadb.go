package pg

import (
	"fmt"
	"time"

	. "data-importer/types"
	"grm-service/common"
	. "grm-service/dbcentral/pg"
	. "grm-service/time"
	"grm-service/util"
)

type MetaDB struct {
	MetaCentralDB
}

// 获取扫描文件列表
func (db MetaDB) GetScanDatas(taskId, fileId string,
	page *common.PageFilter, filter *ResultFilter, count bool) (*ScanResults, error) {
	fucFilter := func(sqlStr string) string {
		if filter == nil {
			return sqlStr
		}
		str := sqlStr
		if len(filter.FileName) > 0 {
			str = fmt.Sprintf(`%s and name ~'%s'`, str, filter.FileName)
		}
		if len(filter.CreateTimeMin) > 0 && len(filter.CreateTimeMax) > 0 {
			str = fmt.Sprintf(`%s and create_time between '%s' and '%s'`, str, filter.CreateTimeMin, filter.CreateTimeMax)
		}
		if len(filter.FileSizeMin) > 0 && len(filter.FileSizeMax) > 0 {
			str = fmt.Sprintf(`%s and file_size between %s and %s`, str, filter.FileSizeMin, filter.FileSizeMax)
		}
		//if filter.ResolutionMax > 0 && filter.ResolutionMax > filter.ResolutionMin {
		//	str = fmt.Sprintf(`%s and resolution between %f and %f`, str, filter.ResolutionMin, filter.ResolutionMax)
		//}
		//if len(filter.RefSystem) > 0 {
		//	str = fmt.Sprintf(`%s and ref_system = '%s'`, str, filter.RefSystem)
		//}
		//if len(filter.Pyramid) > 0 {
		//	str = fmt.Sprintf(`%s and has_pyramid = '%s'`, str, filter.Pyramid)
		//}
		return str
	}
	var results ScanResults
	var total int
	if len(fileId) == 0 && count {
		sql := fmt.Sprintf(`select count(*) from pre_data_object where job_id = '%s'`, taskId)
		sql = fucFilter(sql)
		rows, err := db.Conn.Query(sql)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		if rows.Next() {
			err = rows.Scan(&total)
			if err != nil {
				return nil, err
			}
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}
		results.Total = total
	}

	sqlStr := fmt.Sprintf(`select uuid,name,path,file_size,tags,data_type,sub_type,ref_system,
										create_time from pre_data_object where job_id = '%s'`, taskId)
	if len(fileId) > 0 {
		sqlStr = fmt.Sprintf(`%s and uuid = '%s'`, sqlStr, fileId)
	}
	sqlStr = fucFilter(sqlStr)
	sqlStr = util.PageFilterSql(sqlStr, "uuid", page)
	//fmt.Println(sqlStr)

	rowsData, err := db.Conn.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rowsData.Close()

	var create_time time.Time
	for rowsData.Next() {
		var result ScanResult
		err = rowsData.Scan(&result.FileId, &result.FileName, &result.FilePath, &result.FileSize,
			&result.Tags, &result.FileType, &result.SubType, &result.RefSystem, &create_time)
		if err != nil {
			fmt.Printf("GetScanDatas rows.Scan error: %s\n", err.Error())
			continue
		} else {
			result.CreateTime = GetTimeStd(create_time)
			results.Datas = append(results.Datas, result)
		}
	}

	if err := rowsData.Err(); err != nil {
		return nil, err
	}
	return &results, nil
}

// 更新data url
func (db MetaDB) UpdateDataUrl(data, url string) error {
	sql := fmt.Sprintf(`update data_object set data_url = '%s'  where uuid = '%s'`,
		url, data)
	_, err := db.Conn.Exec(sql)
	return err
}
