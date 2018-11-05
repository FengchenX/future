package pg

import (
	"encoding/json"
	"fmt"
	"time"

	. "data-manager/types"
	"grm-service/common"

	. "grm-service/dbcentral/pg"
	"grm-service/log"
	. "grm-service/time"
	"grm-service/util"
)

type SystemDB struct {
	SysCentralDB
}

// 获取系统类型集合 allTypes:是否包含子类型
func (db SystemDB) GetDataTypes(allTypes bool) (DataTypeList, error) {
	var types DataTypeList

	var sql string
	if allTypes {
		sql = fmt.Sprintf(`select name,label,parent,is_obsoleted,extension,create_time,description from 
						data_type order by name`)
	} else {
		// 只有父类型
		sql = fmt.Sprintf(`select name,label,parent,is_obsoleted,extension,create_time,description from 
						data_type where parent = '' order by name`)
	}

	rows, err := db.Conn.Query(sql)
	if err != nil {
		return types, err
	}
	defer rows.Close()

	for rows.Next() {
		var data DataType
		var create_time time.Time
		err = rows.Scan(&data.Name, &data.Label, &data.Parent, &data.IsObsoleted, &data.Extensions,
			&create_time, &data.Description)
		if err != nil {
			log.Error("GetDataTypes rows.Scan error: %s\n", err.Error())
			continue
		} else {
			data.CreateTime = GetTimeStd(create_time)
			types = append(types, data)
		}
	}
	if err := rows.Err(); err != nil {
		return types, err
	}
	return types, nil
}

// 更新数据类型基本信息
func (db SystemDB) UpdateTypeInfo(info *DataType) error {
	var comma, filter string
	if len(info.Label) > 0 {
		filter = fmt.Sprintf("%s %s label = '%s'", filter, comma, info.Label)
		comma = ","
	}

	if len(info.Description) > 0 {
		filter = fmt.Sprintf("%s %s description = '%s'", filter, comma, info.Description)
		comma = ","
	}

	sql := fmt.Sprintf(`update data_type set %s where name = '%s'`, filter, info.Name)
	fmt.Println(sql)
	_, err := db.Conn.Exec(sql)
	return err
}

// 合并父类元元数据到子类
func mergeMeta(parent, child *Meta) Meta {
	for i, pValue := range parent.DataMeta {
		inGroup := false
		for _, cValue := range child.DataMeta {
			if pValue.Group == cValue.Group {
				parent.DataMeta[i].Values = append(parent.DataMeta[i].Values, cValue.Values...)
				inGroup = true
				break
			}
		}
		if !inGroup {
			parent.DataMeta = append(parent.DataMeta, child.DataMeta...)
		}
	}
	child.DataMeta = parent.DataMeta
	return *child
}

// 获取类型元元信息, 是否包含子类型
func (db SystemDB) GetTypeMetas(dataType string, children bool) (*Meta, error) {
	var meta Meta
	var sql string
	if children {
		sql = fmt.Sprintf(`select name, parent, meta from data_type where name = '%s' 
						or parent = '%s'`, dataType, dataType)
	} else {
		sql = fmt.Sprintf(`select name, parent, meta from data_type where name = '%s'`, dataType)
	}
	log.Info("GetTypeMetas: ", sql)

	rows, err := db.Conn.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name, parent, metaStr string
		err = rows.Scan(&name, &parent, &metaStr)
		if err != nil {
			log.Error("GetTypeMetas rows.Scan error: %s\n", err.Error())
			continue
		}

		// 获取特定类型元元数据
		if name == dataType {
			// 类型
			if err := json.Unmarshal([]byte(metaStr), &meta); err != nil {
				return nil, err
			}
			// 有父类型
			if len(parent) > 0 {
				pmeta, err := db.GetTypeMetas(parent, false)
				if err != nil {
					log.Errorf("Failed to get type(%s)( parent :(%s)) meta: ", dataType, parent, err)
					return nil, err
				}
				meta = mergeMeta(pmeta, &meta)
				break
			}
		} else if parent == dataType {
			// 子类型
			var cmeta Meta
			if err := json.Unmarshal([]byte(metaStr), &cmeta); err != nil {
				return nil, err
			}
			meta.ChildMeta = append(meta.ChildMeta, cmeta)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &meta, nil
}

func (db SystemDB) getMetaFileds(dataType string) (*Meta, error) {
	var meta Meta
	sql := fmt.Sprintf(`select meta from data_type where name = '%s'`, dataType)
	log.Info("GetTypeMetas: ", sql)

	var metaStr string
	err := db.Conn.QueryRow(sql).Scan(&metaStr)
	if err != nil {
		log.Error("getMetaFileds rows.Scan error: %s\n", err.Error())
		return nil, err
	}
	if err := json.Unmarshal([]byte(metaStr), &meta); err != nil {
		return nil, err
	}
	return &meta, nil
}

// 添加元数据信息字段
func (db SystemDB) AddMetaField(typeName string, req *MetaFieldReq) (*Meta, error) {
	meta, err := db.getMetaFileds(typeName)
	if err != nil {
		return nil, err
	}

	field := MetaValue{
		Name:     req.Name,
		Title:    req.Title,
		Type:     req.Type,
		Required: req.Required,
		ReadOnly: req.ReadOnly,
		//Query:    req.Query,
		Classify: req.Classify,
	}

	index := 0
	for ; index < len(meta.DataMeta); index++ {
		if meta.DataMeta[index].Group == req.Group {
			meta.DataMeta[index].Values = append(meta.DataMeta[index].Values, field)
			break
		}
	}
	if index == len(meta.DataMeta) {
		group := GroupMeta{Group: req.Group}
		if req.Group == "Basic Information" {
			group.Label = util.TR("Basic Information")
		}
		group.Values = append(group.Values, field)
		meta.DataMeta = append(meta.DataMeta, group)
	}
	metaStr, err := json.Marshal(meta)
	if err != nil {
		return nil, err
	}
	sql := fmt.Sprintf(`update data_type set meta = '%s' where name = '%s'`, metaStr, typeName)
	fmt.Println(sql)
	_, err = db.Conn.Exec(sql)
	return meta, err
}

// 移除元数据信息字段
func (db SystemDB) DelMetaField(typeName, group, field string) (*Meta, error) {
	meta, err := db.getMetaFileds(typeName)
	if err != nil {
		return nil, err
	}

	for index := 0; index < len(meta.DataMeta); index++ {
		if meta.DataMeta[index].Group == group {
			for i, _ := range meta.DataMeta[index].Values {
				val := meta.DataMeta[index].Values[i]
				if val.Name == field {
					meta.DataMeta[index].Values = append(meta.DataMeta[index].Values[:i], meta.DataMeta[index].Values[i+1:]...)
					break
				}
			}
			break
		}
	}
	metaStr, err := json.Marshal(meta)
	if err != nil {
		return nil, err
	}
	sql := fmt.Sprintf(`update data_type set meta = '%s' where name = '%s'`, metaStr, typeName)
	fmt.Println(sql)
	_, err = db.Conn.Exec(sql)
	return meta, err
}

// 更新元数据信息字段
func (db SystemDB) UpdateMetaField(typeName string, req *MetaFieldReq) (*Meta, error) {
	meta, err := db.getMetaFileds(typeName)
	if err != nil {
		return nil, err
	}

	field := MetaValue{
		Name:     req.Name,
		Title:    req.Title,
		Type:     req.Type,
		Required: req.Required,
		ReadOnly: req.ReadOnly,
		//Query:    req.Query,
		Classify: req.Classify,
	}

	for index := 0; index < len(meta.DataMeta); index++ {
		if meta.DataMeta[index].Group == req.Group {
			for i, _ := range meta.DataMeta[index].Values {
				val := &meta.DataMeta[index].Values[i]
				if val.Name == req.Name {
					meta.DataMeta[index].Values[i] = field
					break
				}
			}
			break
		}
	}
	metaStr, err := json.Marshal(meta)
	if err != nil {
		return nil, err
	}
	sql := fmt.Sprintf(`update data_type set meta = '%s' where name = '%s'`, metaStr, typeName)
	fmt.Println(sql)
	_, err = db.Conn.Exec(sql)
	return meta, err
}

// 添加数据集
func (db SystemDB) AddDataSet(data DataSet) (string, error) {
	if len(data.ID) < 1 {
		data.ID = util.NewUUID()
	}

	sql := fmt.Sprintf(`insert into data_set(id,name,user_id,type,class,description) 
							values ('%s','%s','%s', '%s','%s','%s')`,
		data.ID, data.Name, data.User, data.Type, data.Class, data.Description)
	_, err := db.Conn.Exec(sql)
	if err != nil {
		return "", err
	}
	return data.ID, nil
}

// 获取用户数据集
func (db SystemDB) GetDataSets(userId string) ([]DataSet, error) {
	var datas []DataSet
	var sql string

	if len(userId) > 0 {
		sql = fmt.Sprintf(`select id,name,user_id,create_time,type,class,description from data_set
						 where user_id = '%s' and class = '%s' order by create_time desc`, userId, ClassUeser)
	} else {
		sql = fmt.Sprintf(`select id,name,user_id,create_time,type,class,description from data_set
						 where class = '%s' order by create_time desc`, ClassMarket)
	}

	rows, err := db.Conn.Query(sql)
	if err != nil {
		return datas, err
	}
	defer rows.Close()

	for rows.Next() {
		var data DataSet
		var create_time time.Time
		err = rows.Scan(&data.ID, &data.Name, &data.User, &create_time, &data.Type, &data.Class, &data.Description)
		if err != nil {
			log.Printf("rows.Scan error: %s\n", err.Error())
			return datas, err
		}
		data.CreateTime = GetTimeStd(create_time)
		datas = append(datas, data)
	}
	if err := rows.Err(); err != nil {
		return datas, err
	}
	return datas, nil
}

// 移除数据集
func (db SystemDB) DelDataSet(dataset string) error {
	// 检查数据集下是否有数据
	sql := fmt.Sprintf(`select count(*) from ref_data_user where data_set = '%s'`, dataset)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	var count int
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			log.Printf("DelDataSet rows.Scan error: %s\n", err.Error())
			return err
		}
	}
	if count > 0 {
		return fmt.Errorf(util.TR("Dataset is not empty:%d", count))
	}

	// 移除数据集
	sql = fmt.Sprintf(`delete from data_set where id = '%s'`, dataset)
	_, err = db.Conn.Exec(sql)
	return err
}

// 更新数据集
func (db SystemDB) UpdateDataSet(id, name, description string) error {
	sql := fmt.Sprintf(`update data_set set name = '%s', description = '%s', 
						update_time = current_timestamp where id = '%s'`,
		name, description, id)
	_, err := db.Conn.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

// 清空数据集
func (db SystemDB) TruncateDataSet(dataset string) error {
	sql := fmt.Sprintf(`delete from ref_data_user where data_set = '%s'`, dataset)
	_, err := db.Conn.Exec(sql)
	return err
}

// 移除数据集下数据
func (db SystemDB) DelDataSetData(dataset, data string) error {
	sql := fmt.Sprintf(`delete from ref_data_user where data_set = '%s' and data_id = '%s'`, dataset, data)
	_, err := db.Conn.Exec(sql)
	return err
}

// 添加数据集数据 DataSubscribed
func (db SystemDB) AddDataSetData(dataset, data, user, sourceType string) error {
	sql := fmt.Sprintf(`select class from data_set where id = '%s'`, dataset)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	var class string
	if rows.Next() {
		err = rows.Scan(&class)
		if err != nil {
			log.Printf("AddDataSetData rows.Scan error: %s\n", err.Error())
			return err
		}
	}
	isMarket := "false"
	if class == ClassMarket {
		isMarket = "true"
	}
	sql = fmt.Sprintf(`insert into ref_data_user(data_id,user_id,data_set,type,is_market)
						values('%s','%s','%s','%s',%s)`, data, user, dataset, sourceType, isMarket)
	_, err = db.Conn.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

// 批量添加数据集数据 DataSubscribed
func (db SystemDB) AddDataSetDatas(dataset, user, sourceType, srcDataSet string) error {
	sql := fmt.Sprintf(`select class from data_set where id = '%s'`, dataset)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	var class string
	if rows.Next() {
		err = rows.Scan(&class)
		if err != nil {
			log.Printf("AddDataSetData rows.Scan error: %s\n", err.Error())
			return err
		}
	}
	isMarket := "false"
	if class == ClassMarket {
		isMarket = "true"
	}

	sql = fmt.Sprintf(`insert into ref_data_user(data_id,user_id,data_set,type,is_market) 
						select data_id,'%s','%s','%s',%s from ref_data_user 
						where data_set = '%s'`, user, dataset, sourceType, isMarket, srcDataSet)
	_, err = db.Conn.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

// 获取用户数据图层列表
func (db SystemDB) GetDataLayers(data, user, dataName string) ([]*common.DataLayer, error) {
	var layers []*common.DataLayer
	sql := fmt.Sprintf(`select id,name,style,description,create_time,is_default,srs,wms,wmts,wfs,wms_pub,wmts_pub,
						snapshot from ref_data_layer where data_id = '%s' and (user_id = '%s' or is_default = true) 
						order by create_time asc`, data, user)
	//fmt.Println(sql)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return layers, err
	}
	defer rows.Close()

	for rows.Next() {
		var layer common.DataLayer
		layer.Data = data
		layer.User = user
		layer.DataName = dataName
		var create_time time.Time
		err = rows.Scan(&layer.Layer, &layer.Name, &layer.Style, &layer.Description, &create_time, &layer.IsDefault,
			&layer.Srs, &layer.WMS, &layer.Wmts, &layer.Wfs, &layer.WmsUrl, &layer.WmtsUrl, &layer.SnapShot)
		if err != nil {
			log.Printf("GetDataLayers rows.Scan error: %s\n", err.Error())
			return layers, err
		}
		layer.CreateTime = GetTimeStd(create_time)
		layers = append(layers, &layer)
	}
	if err := rows.Err(); err != nil {
		return layers, err
	}
	return layers, nil
}

// 获取图层信息
func (db SystemDB) GetDataLayer(data, layerId string) (*common.DataLayer, error) {
	var layer common.DataLayer
	layer.Data = data
	layer.Layer = layerId

	sql := fmt.Sprintf(`select id,user_id,name,style,description,create_time,is_default,srs,wms,wmts,wfs,wms_pub,wmts_pub,
						snapshot from ref_data_layer where data_id = '%s' and id = '%s'`, data, layerId)
	//fmt.Println(sql)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var create_time time.Time
		err = rows.Scan(&layer.Layer, &layer.User, &layer.Name, &layer.Style, &layer.Description, &create_time, &layer.IsDefault,
			&layer.Srs, &layer.WMS, &layer.Wmts, &layer.Wfs, &layer.WmsUrl, &layer.WmtsUrl, &layer.SnapShot)
		if err != nil {
			log.Printf("GetDataLayers rows.Scan error: %s\n", err.Error())
			return nil, err
		}
		layer.CreateTime = GetTimeStd(create_time)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &layer, nil
}

// 移除数据图层
func (db SystemDB) DelDataLayer(layer, user string) error {
	sql := fmt.Sprintf(`delete from ref_data_layer where id = '%s' and user_id = '%s'`, layer, user)
	_, err := db.Conn.Exec(sql)
	return err
}

// 更新图层快视图
func (db SystemDB) UpdateLayerSnap(layer, url string) error {
	sql := fmt.Sprintf(`update ref_data_layer set snapshot = '%s'  where id = '%s'`, url, layer)
	_, err := db.Conn.Exec(sql)
	return err
}

// 编辑数据图层
func (db SystemDB) UpdateDataLayer(data, layer, user string, args *UpdateLayerReq) error {
	if args.IsDefault {
		sql := fmt.Sprintf(`update ref_data_layer set is_default = false where data_id = '%s' and user_id = '%s'`, data, user)
		fmt.Println(sql)
		if _, err := db.Conn.Exec(sql); err != nil {
			return err
		}
	}

	var comma, filter string
	if len(args.Name) > 0 {
		filter = fmt.Sprintf("%s %s name = '%s'", filter, comma, args.Name)
		comma = ","
	}

	if len(args.Style) > 0 {
		filter = fmt.Sprintf("%s %s style = '%s'", filter, comma, args.Style)
		comma = ","
	}
	if len(args.Description) > 0 {
		filter = fmt.Sprintf("%s %s description = '%s'", filter, comma, args.Description)
		comma = ","
	}

	if args.IsDefault {
		filter = fmt.Sprintf("%s %s is_default = true", filter, comma)
		comma = ","
	}

	sql := fmt.Sprintf(`update ref_data_layer set %s  where id = '%s' and user_id = '%s' and data_id = '%s'`,
		filter, layer, user, data)
	//fmt.Println(sql)
	_, err := db.Conn.Exec(sql)
	return err
}

// 添加图层样式
func (db SystemDB) AddStyle(style *common.LayerStyle) (*common.LayerStyle, error) {
	var createTime time.Time
	sql := fmt.Sprintf(`insert into ref_layer_style(name,type,user_id,description) 
						values('%s','%s','%s','%s') returning id,create_time`,
		style.Name, style.Type, style.User, style.Description)
	if err := db.Conn.QueryRow(sql).Scan(&style.Id, &createTime); err != nil {
		return nil, err
	}
	style.CreateTime = GetTimeStd(createTime)
	return style, nil
}

// 编辑样式
func (db SystemDB) UpdateStyle(style *common.LayerStyle) error {
	var comma, filter string
	if len(style.Name) > 0 {
		filter = fmt.Sprintf("%s %s name = '%s'", filter, comma, style.Name)
		comma = ","
	}

	if len(style.Description) > 0 {
		filter = fmt.Sprintf("%s %s description = '%s'", filter, comma, style.Description)
		comma = ","
	}

	sql := fmt.Sprintf(`update ref_layer_style set %s  where id = %s and user_id = '%s'`,
		filter, style.Id, style.User)
	_, err := db.Conn.Exec(sql)
	return err
}

// 获取用户图层
func (db SystemDB) GetStyles(user, styleType string) ([]*common.LayerStyle, error) {
	var styles []*common.LayerStyle

	var sql string
	if len(styleType) > 0 {
		sql = fmt.Sprintf(`select id,name,type,description,create_time,user_id from ref_layer_style where type = '%s' and (user_id = '%s' 
						or user_id = '') order by create_time asc`, styleType, user)
	} else {
		sql = fmt.Sprintf(`select id,name,type,description,create_time,user_id from ref_layer_style where user_id = '%s' 
						or user_id = '' order by create_time asc`, user)
	}

	//fmt.Println(sql)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return styles, err
	}
	defer rows.Close()

	for rows.Next() {
		var style common.LayerStyle
		var create_time time.Time
		err = rows.Scan(&style.Id, &style.Name, &style.Type, &style.Description, &create_time, &style.User)
		if err != nil {
			log.Printf("GetStyles rows.Scan error: %s\n", err.Error())
			return styles, err
		}
		style.CreateTime = GetTimeStd(create_time)
		styles = append(styles, &style)
	}
	if err := rows.Err(); err != nil {
		return styles, err
	}
	return styles, nil
}

// 获取样式
func (db SystemDB) GetStyle(styleId string) (*common.LayerStyle, error) {
	var style common.LayerStyle
	style.Id = styleId

	var create_time time.Time
	sql := fmt.Sprintf(`select id,name,type,description,create_time from ref_layer_style where id = %s `, styleId)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&style.Id, &style.Name, &style.Type, &style.Description, &create_time)
		if err != nil {
			log.Printf("GetStyle rows.Scan error: %s\n", err.Error())
			return nil, err
		}
		style.CreateTime = GetTimeStd(create_time)
	}
	return &style, nil
}

// 移除样式
func (db SystemDB) DelStyle(style, user string) (string, error) {
	if len(user) == 0 {
		return "", nil
	}
	var name string
	sql := fmt.Sprintf(`delete from ref_layer_style where id = %s and user_id = '%s' returning name`, style, user)
	if err := db.Conn.QueryRow(sql).Scan(&name); err != nil {
		return "", err
	}
	return name, nil
}
