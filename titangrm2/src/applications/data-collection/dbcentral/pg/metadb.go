package pg

import (
	"encoding/json"
	"fmt"
	"strings"

	. "grm-service/dbcentral/pg"

	"applications/data-collection/types"
)

type MetaDB struct {
	MetaCentralDB
}

// 添加数据记录
func (db MetaDB) AddDataObject(data, name, dataset, device, user, shpType string) error {
	metajson := fmt.Sprintf(`
		{
    "full_valid" : false,
    "label" : "矢量数据",
    "metadata" : [
       {
          "group" : "Basic Information",
          "label" : "基本信息",
          "value" : [
             {
                "classify" : "false",
                "name" : "create_time",
                "required" : "false",
                "system" : "false",
                "title" : "创建时间",
                "type" : "time",
                "value" : "1000-01-01"
             },
             {
                "classify" : "true",
                "name" : "data_time",
                "required" : "true",
                "system" : "false",
                "title" : "数据时间",
                "type" : "date",
                "value" : "1000-01-01"
             },
             {
                "classify" : "false",
                "name" : "description",
                "required" : "false",
                "system" : "false",
                "title" : "描述",
                "type" : "string",
                "value" : ""
             },
             {
                "classify" : "false",
                "name" : "envelope",
                "required" : "true",
                "system" : "false",
                "title" : "地理坐标",
                "type" : "array",
                "value" : ""
             },
             {
                "classify" : "false",
                "name" : "feature_class_count",
                "required" : "true",
                "system" : "true",
                "title" : "类型数量",
                "type" : "int",
                "value" : ""
             },
             {
                "classify" : "false",
                "name" : "feature_nums",
                "required" : "true",
                "system" : "true",
                "title" : "要素个数",
                "type" : "bigint",
                "value" : ""
             },
             {
                "classify" : "false",
                "name" : "field_list",
                "required" : "false",
                "system" : "true",
                "title" : "字段列表",
                "type" : "string",
                "value" : ""
             },
             {
                "classify" : "false",
                "name" : "file_size",
                "required" : "false",
                "system" : "true",
                "title" : "文件大小",
                "type" : "bigint",
                "value" : ""
             },
             {
                "classify" : "false",
                "name" : "name",
                "required" : "true",
                "system" : "false",
                "title" : "矢量名称",
                "type" : "string",
                "value" : "%s"
             },
             {
                "classify" : "false",
                "name" : "north_east_x",
                "required" : "true",
                "system" : "true",
                "title" : "东北图廓角点X坐标",
                "type" : "int",
                "value" : ""
             },
             {
                "classify" : "false",
                "name" : "north_east_y",
                "required" : "true",
                "system" : "true",
                "title" : "东北图廓角点Y坐标",
                "type" : "int",
                "value" : ""
             },
             {
                "classify" : "false",
                "name" : "north_west_x",
                "required" : "true",
                "system" : "true",
                "title" : "西北图廓角点X坐标",
                "type" : "int",
                "value" : ""
             },
             {
                "classify" : "false",
                "name" : "north_west_y",
                "required" : "true",
                "system" : "true",
                "title" : "西北图廓角点Y坐标",
                "type" : "int",
                "value" : ""
             },
             {
                "classify" : "false",
                "name" : "path",
                "required" : "true",
                "system" : "true",
                "title" : "路径",
                "type" : "string",
                "value" : ""
             },
             {
                "classify" : "false",
                "name" : "ref_system",
                "required" : "true",
                "system" : "false",
                "title" : "参考系",
                "type" : "string",
                "value" : ""
             },
             {
                "classify" : "true",
                "name" : "shp_type",
                "required" : "true",
                "system" : "true",
                "title" : "矢量类型",
                "type" : "string",
                "value" : ""
             },
             {
                "classify" : "false",
                "name" : "south_east_x",
                "required" : "true",
                "system" : "true",
                "title" : "东南图廓角点X坐标",
                "type" : "int",
                "value" : ""
             },
             {
                "classify" : "false",
                "name" : "south_east_y",
                "required" : "true",
                "system" : "true",
                "title" : "东南图廓角点Y坐标",
                "type" : "int",
                "value" : ""
             },
             {
                "classify" : "false",
                "name" : "south_west_x",
                "required" : "true",
                "system" : "true",
                "title" : "西南图廓角点X坐标",
                "type" : "int",
                "value" : ""
             },
             {
                "classify" : "false",
                "name" : "south_west_y",
                "required" : "true",
                "system" : "true",
                "title" : "西南图廓角点Y坐标",
                "type" : "int",
                "value" : ""
             }
          ]
       }
    ],
    "type" : "Shape"
 }`, name)

	envelope := `ST_GeomFromGeoJSON('{
	"type": "Polygon",
	"coordinates": [
		[
			[116.28852713281697, 40.22718090480737],
			[116.87148916894972, 40.22980208808957],
			[116.83097708398881, 41.06531175295575],
			[116.25556814844191, 41.06065225097328],
			[116.28852713281697, 40.22718090480737]
		]
	]
	}')`
	sql := fmt.Sprintf(`insert into meta_object(uuid,name,meta_type,create_time,dataset,device,load_user,
						data_time,ref_system,shp_type,path,file_size,meta_json,envelope) 
						values('%s','%s','Shape',current_timestamp,'%s','%s',
						'%s','1000-01-01','GCS_WGS_1984','%s','',-1,'%s',%s)`,
		data, name, dataset, device, user, shpType, metajson, envelope)
	_, err := db.Conn.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

// 修改元数据信息
func (db *MetaDB) UpdateDataMeta(dataId string, req *types.UpdateMetaRequest) error {
	columns := `name,path,file_size,create_time,data_time,
								projection_type,resolution,size,ref_system,
								coord_unit,thumb_path,snap_path,
								shp_type, feature_nums, sat_type, sensor, description`

	sql := fmt.Sprintf(`select meta_json from meta_object where uuid = '%s'`, dataId)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	var meta_json string
	if rows.Next() {
		err = rows.Scan(&meta_json)
		if err != nil {
			fmt.Printf("rows.Scan error: %s\n", err.Error())
			return err
		} else {
			var metas types.DataMeta
			err := json.Unmarshal([]byte(meta_json), &metas)
			if err != nil {
				fmt.Printf("Failed to parse meta json : %s\n", err.Error())
				return err
			}
			for iGroup, _ := range metas.MetaData {
				if len(req.Group) > 0 && metas.MetaData[iGroup].Group != req.Group {
					continue
				}

				for _, meta := range req.Metas {
					exists := false
					for jData, _ := range metas.MetaData[iGroup].Value {
						if metas.MetaData[iGroup].Value[jData].Name == meta.Name {
							exists = true

							if metas.MetaData[iGroup].Value[jData].Value != meta.Value {
								metas.MetaData[iGroup].Value[jData].Value = meta.Value
								dataModify := metas.MetaData[iGroup].Value[jData].Modified
								modify, ok := dataModify.(bool)
								if ok && !modify {
									metas.MetaData[iGroup].Value[jData].Modified = true
								} else {
									modify, ok := dataModify.(string)
									if ok && modify != "true" {
										metas.MetaData[iGroup].Value[jData].Modified = true
									}
								}

							}
							break
						}
					}
					if !exists {
						metas.MetaData[iGroup].Value = append(metas.MetaData[iGroup].Value, meta)
					}
				}

				//				for jData, _ := range metas.MetaData[iGroup].Value {
				//					for _, meta := range req.Metas {
				//						if metas.MetaData[iGroup].Value[jData].Name == meta.Name {
				//							metas.MetaData[iGroup].Value[jData].Value = meta.Value
				//						}
				//					}
				//				}
			}

			ret, err := json.Marshal(metas)
			if err != nil {
				return err
			}
			// 更新meta表
			sql := fmt.Sprintf(`update meta_object set meta_json = '%s'`, string(ret))
			for _, meta := range req.Metas {
				if strings.LastIndex(columns, meta.Name) != -1 {
					sql = fmt.Sprintf(`%s, %s = '%s'`, sql, meta.Name, meta.Value)
				}
			}

			//if req.Tags != base.NullTags {
			//	sql = fmt.Sprintf(`%s, tags = '%s'`, sql, req.Tags)
			//}
			if req.DisplayField != "" {
				sql = fmt.Sprintf(`%s, display_field = '%s'`, sql, req.DisplayField)
			}

			sql = fmt.Sprintf(`%s where uuid = '%s'`, sql, dataId)
			_, err = db.Conn.Exec(sql)
			if err != nil {
				return err
			}
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

// 移除数据调查
func (db *MetaDB) DelDataObject(colId string) error {
	sql := fmt.Sprintf(`delete from meta_object where uuid = '%s'`, colId)
	_, err := db.Conn.Exec(sql)
	return err
}
