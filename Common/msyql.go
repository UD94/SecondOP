package Common

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/go-ini/ini"
)

//注意方法名大写，就是public
func InitDB(DB *sql.DB) error {

	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("ini file not found!")
	}
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{cfg.Section("mysql").Key("user").String(), ":", cfg.Section("mysql").Key("password").String(), "@tcp(", cfg.Section("mysql").Key("ip").String(), ":", cfg.Section("mysql").Key("port").String(), ")/", cfg.Section("mysql").Key("database").String(), "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"

	DB, err = sql.Open("mysql", path)
	//设置数据库最大连接数
	if err != nil {
		fmt.Println("opon database fail")
		return errors.New("database no ready")

	}
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)

	fmt.Println("connnect success")

	return nil

}
