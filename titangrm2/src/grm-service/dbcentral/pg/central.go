package pg

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
)

const (
	AuthDBName = "TitanCloud.Auth"
	SysDBName  = "TitanCloud.System"
	MetaDBName = "TitanCloud.Meta"
	DataDBName = "TitanCloud.Data"
)

type ConnConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

type Central struct {
	Config *ConnConfig
	Conn   *sql.DB
}

func ConnectDB(host, dbName, user, password string) (Central, error) {
	//db, err := sql.Open("pgx", "postgres://postgres:123456@192.168.1.189/TitanCloud.System?sslmode=disable")
	dataSource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbName)
	db, err := sql.Open("pgx", dataSource)
	if err != nil {
		return Central{}, err
	}
	return Central{
		Config: &ConnConfig{
			Host:     host,
			Database: dbName,
			User:     user,
			Password: password,
		},
		Conn: db,
	}, nil
}

func ConnectDBUrl(dataSource string) (Central, error) {
	//db, err := sql.Open("pgx", "postgres://postgres:123456@192.168.1.189/TitanCloud.System?sslmode=disable")
	//	dataSource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbName)
	db, err := sql.Open("pgx", dataSource)
	if err != nil {
		return Central{}, err
	}
	return Central{
		Conn: db,
	}, nil
}

func (db *Central) DisConnect() error {
	return db.Conn.Close()
}
