package Function

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/UD94/SecondOP/Common"
)

type MD5Struct struct {
	NTLM     string
	Password string
}

func Md5_query(hash_str string) (string, error) {
	var user MD5Struct

	DB := new(sql.DB)
	defer DB.Close()
	status := Common.InitDB(DB)
	if status == nil {
		err := DB.QueryRow("SELECT * FROM md5 WHERE ntlm = ?", hash_str).Scan(&user.NTLM, &user.Password)
		if err != nil {
			fmt.Println("查询出错了")
			return "nopass", errors.New("nopass")
		}
		DB.Close()
		return user.Password, nil
	} else {
		return "nodatabase", errors.New("no database")
	}

}
func MD5_insert(hash_str string, password string) (string, error) {
	DB := new(sql.DB)
	defer DB.Close()
	err := Common.InitDB(DB)
	if err == nil {
		_, err := DB.Exec("insert into md5(ntlm,password) values(?,?)", hash_str, password)
		if err != nil {
			fmt.Println("新增数据错误", err)
			return "inserterror", errors.New("insert error")
		}
	} else {
		return "nodatabase", errors.New("no database")
	}
	return "success", nil
}
