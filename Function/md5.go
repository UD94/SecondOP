package Function

import (
	"database/sql"
	"errors"
	"fmt"
)

type MD5Struct struct {
	Hash     string
	Password string
}

func Md5_query(DB *sql.DB, hash_str string) (string, error) {
	var user MD5Struct

	err := DB.QueryRow("SELECT * FROM ntlm WHERE hash = ?", hash_str).Scan(&user.Hash, &user.Password)
	if err != nil {
		fmt.Println("查询出错了")

		return "nopass", errors.New("no pass")
	}

	return user.Password, nil

}
func MD5_insert(DB *sql.DB, hash_str string, password string) (string, error) {

	_, err := DB.Exec("insert into ntlm(hash,password) values(?,?)", hash_str, password)
	if err != nil {
		fmt.Println("新增数据错误", err)
		return "inserterror", errors.New("insert error")
	}

	return "success", nil
}
