//author xinbing
//time 2018/9/4 17:55
package db

import "github.com/pkg/errors"

type DBConfig struct {
	DBAddr	string
	AutoCreateTables []interface{} //自动创建的表
	MaxIdleConns int
	MaxOpenConns int
	LogMode		 bool
}

func (p *DBConfig) check() error {
	if p.DBAddr == "" {
		return errors.New("empty sql addr")
	}
	if p.MaxIdleConns <= 0 {
		p.MaxIdleConns = 10
	}
	if p.MaxOpenConns <= 0 {
		p.MaxOpenConns = 100
	}
	return nil
}