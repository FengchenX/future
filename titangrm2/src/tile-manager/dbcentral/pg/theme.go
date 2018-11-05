package pg

import (
	dbsql "database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx"
	"grm-service/log"
	"grm-service/util"
	"strings"
	"tile-manager/types"
	"time"
)

var (
	ErrThemeExists = errors.New("Theme is already exists")
)

func (db SystemDB) AddTheme(theme *types.Theme) (int, error) {
	sql := fmt.Sprintf(`insert into theme(name,image_type,tile_size,tile_format,
							srs,no_data,time_resolution,user_id,
							create_time,description,transparency,device,projection) 
							values ('%s','%s',%d,'%s','%s',%f,'%s','%s',%s,'%s','%s', %d,'%s') returning id`,
		theme.Name, theme.ImageType, theme.TileSize, theme.TileFormat,
		theme.Srs, theme.NoData, theme.TimeResolution, theme.AuthUser,
		util.GetTimeNowDB(), theme.Description, theme.Transparency, theme.DeviceId, theme.Projection)
	stmt, err := db.Conn.Prepare(sql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRow().Scan(&id)
	if err != nil {
		if pgErr, ok := err.(pgx.PgError); ok {
			if pgErr.Code == "23505" {
				return -1, ErrThemeExists
			}
		}
		return -1, err
	}
	return id, nil
}

func (db SystemDB) UpdateTheme(theme *types.Theme) error {
	var comma string
	sql := fmt.Sprintf(`update theme set`)
	if len(theme.Name) > 0 {
		sql = fmt.Sprintf(`%s name = '%s'`, sql, theme.Name)
		comma = ","
	}
	if len(theme.ImageType) > 0 {
		sql = fmt.Sprintf(`%s %s image_type = '%s'`, sql, comma, theme.ImageType)
		comma = ","
	}
	if theme.TileSize > 0 {
		sql = fmt.Sprintf(`%s %s tile_size = %d`, sql, comma, theme.TileSize)
		comma = ","
	}
	if len(theme.TileFormat) > 0 {
		sql = fmt.Sprintf(`%s %s tile_format = '%s'`, sql, comma, theme.TileFormat)
		comma = ","
	}
	if len(theme.Srs) > 0 {
		sql = fmt.Sprintf(`%s %s srs = '%s'`, sql, comma, theme.Srs)
		comma = ","
	}
	if theme.NoData != -1 {
		sql = fmt.Sprintf(`%s %s no_data = %f`, sql, comma, theme.NoData)
		comma = ","
	}
	if len(theme.TimeResolution) > 0 {
		sql = fmt.Sprintf(`%s %s time_resolution = '%s'`, sql, comma, theme.TimeResolution)
		comma = ","
	}
	if len(theme.Description) > 0 {
		sql = fmt.Sprintf(`%s %s description = '%s'`, sql, comma, theme.Description)
		comma = ","
	}
	if len(theme.Transparency) > 0 {
		sql = fmt.Sprintf(`%s %s transparency = '%s'`, sql, comma, theme.Transparency)
		comma = ","
	}

	if len(theme.Projection) > 0 {
		sql = fmt.Sprintf(`%s %s projection = '%s'`, sql, comma, theme.Projection)
		comma = ","
	}
	sql = fmt.Sprintf(`%s where id = %d`, sql, theme.Id)

	_, err := db.Conn.Exec(sql)
	return err
}

func (db SystemDB) DelTheme(themeId int, userId string) error {
	sql := fmt.Sprintf(`delete form theme where id = %d`, themeId)
	if len(userId) > 0 {
		sql = fmt.Sprintf(`%s and user_id = '%s'`, sql, userId)
	}
	_, err := db.Conn.Exec(sql)
	return err
}
func (db SystemDB) GetThemes(users []string, theme int, limit, offset int, sort, order string) (int, []types.Theme, error) {
	var total int
	var themes []types.Theme
	if theme > 0 {
		total = 1
	} else {
		sql := fmt.Sprintf(`select count(*) from theme where`)
		args := ""
		for _, user := range users {
			args = args + fmt.Sprintf("or user_id = '%s'", user)
		}
		args = strings.TrimPrefix(args, "or")
		sql = sql + args

		rows, err := db.Conn.Query(sql)
		if err != nil {
			return 0, themes, err
		}
		defer rows.Close()

		if rows.Next() {
			err = rows.Scan(&total)
			if err != nil {
				log.Errorf("rows.Scan error: %s\n", err.Error())
				return 0, themes, err
			}
		}
		if err := rows.Err(); err != nil {
			return 0, themes, err
		}
	}
	sql := fmt.Sprintf(`select id,name,image_type,tile_size,tile_format,
							srs,no_data,time_resolution,user_id,create_time,description,transparency,thumb,device,projection from theme where`)

	args := ""
	for _, user := range users {
		args = args + fmt.Sprintf("or user_id = '%s'", user)
	}
	args = strings.TrimPrefix(args, "or")
	sql = sql + args
	if theme > 0 {
		sql = fmt.Sprintf("%s and id = %d", sql, theme)
	}
	if sort != "" && order != "" {
		sql = fmt.Sprintf(`%s order by %s %s,user_id %s`, sql, sort, order, order)
	}
	if offset >= 0 && limit > 0 {
		sql = fmt.Sprintf(`%s limit %d offset %d`, sql, limit, offset)
	}

	rows, err := db.Conn.Query(sql)
	if err != nil {
		return 0, themes, err
	}
	defer rows.Close()

	var name, image_type, tile_format, srs, time_resolution string
	var user_id, description, transparency, thumb string
	var create_time time.Time
	var id, device int
	var tile_size uint32
	var no_data float32
	var projection dbsql.NullString

	for rows.Next() {
		err = rows.Scan(&id, &name, &image_type, &tile_size, &tile_format,
			&srs, &no_data, &time_resolution, &user_id, &create_time, &description, &transparency,
			&thumb, &device, &projection)
		if err != nil {
			log.Errorf("rows.Scan error: %s\n", err.Error())
			continue
		} else {
			//获取下theme下面的geometry
			sqlWkt := fmt.Sprintf(`SELECT ST_AsText(ST_centroid(geo_wkt)) from patch where theme_id = %d;`, id)
			var center string
			var geoWkt dbsql.NullString
			rowsWkt, err := db.Conn.Query(sqlWkt)
			if err == nil {
				for rowsWkt.Next() {
					rowsWkt.Scan(&geoWkt)
					if geoWkt.Valid {
						center = geoWkt.String[6 : len(geoWkt.String)-1]
						break
					}
				}
				rowsWkt.Close()
			}

			data := types.Theme{
				Id:             id,
				Name:           name,
				ImageType:      image_type,
				TileSize:       tile_size,
				TileFormat:     tile_format,
				Srs:            srs,
				NoData:         no_data,
				TimeResolution: time_resolution,
				AuthUser:       user_id,
				CreateTime:     util.GetTimeStd(create_time),
				Description:    description,
				Transparency:   transparency,
				Thumb:          thumb,
				DeviceId:       device,
				DeviceInfo:     types.Device{},
				Projection:     projection.String,
				Center:         center,
				Session:        "",
			}
			themes = append(themes, data)
		}
	}
	if err := rows.Err(); err != nil {
		return 0, themes, err
	}
	return total, themes, nil
}

func (db SystemDB) UpdateThemePic(id int, thumb string) error {
	sql := fmt.Sprintf(
		`UPDATE theme 
				SET thumb = '%s' 
				WHERE
				"id" =%d`, thumb, id)
	_, err := db.Conn.Exec(sql)
	return err
}
