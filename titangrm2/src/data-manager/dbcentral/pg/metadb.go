package pg

import (
	"fmt"
	"time"

	"data-manager/types"
	"grm-service/common"
	. "grm-service/dbcentral/pg"
	"grm-service/log"
	. "grm-service/time"
	"grm-service/util"
)

type MetaDB struct {
	MetaCentralDB
}

// 获取数据类型
func (db MetaDB) GetDataType(data string) (string, error) {
	sql := fmt.Sprintf(`select data_type,sub_type from data_object where uuid = '%s'`, data)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var dataType, subType string
	if rows.Next() {
		err = rows.Scan(&dataType, &subType)
		if err != nil {
			return "", err
		}
	}
	if err := rows.Err(); err != nil {
		return "", err
	}

	if len(subType) == 0 {
		return dataType, nil
	} else {
		return subType, nil
	}
}

// 更新数据快视图
func (db MetaDB) UpdateDataSnap(data, url string) error {
	sql := fmt.Sprintf(`update data_object set snapshot = '%s'  where uuid = '%s'`, url, data)
	_, err := db.Conn.Exec(sql)
	return err
}

// 更新浏览次数
func (db MetaDB) UpdateDataViewCnt(data string) (int, error) {
	sql := fmt.Sprintf(`update data_object set viewcnt = viewcnt + 1  where uuid = '%s' returning viewcnt`, data)

	var cnt int
	err := db.Conn.QueryRow(sql).Scan(&cnt)
	return cnt, err
}

// 更新数据状态
func (db MetaDB) UpdateDataStatus(data, status string) error {
	sql := fmt.Sprintf(`update data_object set status = '%s'  where uuid = '%s'`, status, data)
	_, err := db.Conn.Exec(sql)
	return err
}

// 更新数据信息
func (db MetaDB) UpdateDataInfo(data string, args *types.UpdateDataInfoReq) error {
	var comma, filter string
	if args.Abstract != common.OmitArg {
		filter = fmt.Sprintf("%s %s abstract = '%s'", filter, comma, args.Abstract)
		comma = ","
	}
	if args.Description != common.OmitArg {
		filter = fmt.Sprintf("%s %s description = '%s'", filter, comma, args.Description)
		comma = ","
	}
	if args.Tags != common.OmitArg {
		filter = fmt.Sprintf("%s %s tags = '%s'", filter, comma, args.Tags)
		comma = ","
	}
	sql := fmt.Sprintf(`update data_object set %s where uuid = '%s'`, filter, data)
	_, err := db.Conn.Exec(sql)
	return err
}

// 获取数据信息
func (db MetaDB) GetDataInfo(data string) (*types.DataInfo, error) {
	sql := fmt.Sprintf(`select name,owner,path,file_size,data_type,sub_type,data_time,tags,description,
							abstract,status,thumb,snapshot,viewcnt,create_time,data_url,storage,
							COALESCE(delete_time,'1970-01-01') from data_object where uuid = '%s'`, data)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var info types.DataInfo
	info.DataId = data
	if rows.Next() {
		var dataTime, createTime, delTime time.Time
		err = rows.Scan(&info.DataName, &info.Owner.Id, &info.Path, &info.Size, &info.DataType, &info.SubType,
			&dataTime, &info.Tags, &info.Description, &info.Abstract, &info.Status, &info.Thumb, &info.SnapShot,
			&info.ViewCnt, &createTime, &info.DataUrl, &info.Storage, &delTime)
		if err != nil {
			return nil, err
		}
		info.DataTime = GetDateStd(dataTime)
		if info.DataTime == "1000-01-01" {
			info.DataTime = ""
		}
		info.CreateTime = GetTimeStd(createTime)
		info.DeleteTime = GetTimeStd(delTime)
		if info.DeleteTime == "1970-01-01 00:00:00" {
			info.DeleteTime = ""
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &info, nil
}

// 获取数据子对象
func (db MetaDB) GetDataSubObjs(dataId string) (*types.DataSubObj, error) {
	var datas types.DataSubObj
	datas.Group = make(map[string]*types.SubGroup)
	sql := fmt.Sprintf(`select id,groups,object_id,file from data_subobject where data_id = '%s'`, dataId)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var group, id, object_id, file string
		err = rows.Scan(&id, &group, &object_id, &file)
		if err != nil {
			log.Error("GetDataSubObjs rows.Scan error: %s\n", err.Error())
			continue
		} else {
			if group == "images" {
				group = util.TR("images")
			} else if group == "common" {
				group = util.TR("no group")
			}

			subGroup, ok := datas.Group[group]
			if ok && subGroup != nil {
				subGroup.Datas[id] = file
				subGroup.Count++
			} else {
				subgroup := types.SubGroup{Datas: map[string]string{id: file}}
				subgroup.Count++
				datas.Group[group] = &subgroup
			}
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &datas, nil
}
