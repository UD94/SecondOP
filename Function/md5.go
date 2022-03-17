package Function

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/UD94/SecondOP/Common"
)

type MD5Struct struct {
	Hash     string
	Password string
}

func Testmysql() {
	var user MD5Struct
	var DB *sql.DB

	DB, status := Common.InitDB("ntlm")
	defer DB.Close()
	if status == nil {
		fmt.Printf("success")
		sqlStr := "SELECT * FROM ntlm WHERE Hash = ?"
		DB.QueryRow(sqlStr, "8e8efa6ef8e8bd8f7469a1e0b5a77dd8").Scan(&user.Hash, &user.Password)
		fmt.Println(user.Password)
	}

}

func Md5_query(hash_str string) (string, error) {
	var user MD5Struct

	var DB *sql.DB

	DB, status := Common.InitDB("ntlm")

	if status == nil {
		err := DB.QueryRow("SELECT * FROM ntlm WHERE hash = ?", hash_str).Scan(&user.Hash, &user.Password)
		if err != nil {
			fmt.Println("查询出错了")
			return "nopass", errors.New("no pass")
		}
		defer DB.Close()
		return user.Password, nil
	} else {
		defer DB.Close()
		return "nodatabase", errors.New("no database")
	}

}
func MD5_insert(hash_str string, password string) (string, error) {
	var DB *sql.DB

	DB, err := Common.InitDB("ntlm")

	if err == nil {
		_, err := DB.Exec("insert into ntlm(hash,password) values(?,?)", hash_str, password)
		if err != nil {
			fmt.Println("新增数据错误", err)
			defer DB.Close()
			return "inserterror", errors.New("insert error")
		}
	} else {
		defer DB.Close()
		return "nodatabase", errors.New("no database")
	}
	defer DB.Close()
	return "success", nil
}
