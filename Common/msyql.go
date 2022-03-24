package Common

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//注意方法名大写，就是public
func InitDB(database string) (*sql.DB, error) {
	var DB *sql.DB

	x := Loadini("config.ini")
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{x["mysql"]["user"], ":", x["mysql"]["password"], "@tcp(", x["mysql"]["ip"], ":", x["mysql"]["port"], ")/", database}, "")

	DB, err := sql.Open("mysql", path)
	//设置数据库最大连接数
	if err != nil {
		fmt.Println("opon database fail")
		return DB, errors.New("database no ready")

	}
	DB.SetMaxOpenConns(500)
	DB.SetMaxIdleConns(500)              //空闲连接200
	DB.SetConnMaxLifetime(time.Hour * 1) //长连接有效期一个小时

	fmt.Println("connnect success")

	return DB, nil

}
