
package  main

import (
    "fmt"
    "github.com/BurntSushi/toml"
)

//Config 订制配置文件解析载体
type Config struct{
    Database *Database
    SQL *SQL
}

//Database 订制Database块
type Database struct {
    Driver    string
    Username  string `toml:"us"` //表示该属性对应toml里的us
    Password string
}
//SQL 订制SQL语句结构
type SQL struct{
    SQL1 string `toml:"sql_1"`
    SQL2 string `toml:"sql_2"`
    SQL3 string `toml:"sql_3"`
    SQL4 string `toml:"sql_4"`
}

var config =new(Config)
func init(){
    //读取配置文件
    _, err := toml.DecodeFile("test.toml",config)
    if err!=nil{
        fmt.Println(err)
    }
}
func main() {
      fmt.Println(config)
      fmt.Println(config.Database)
      fmt.Println(config.SQL.SQL1)
}