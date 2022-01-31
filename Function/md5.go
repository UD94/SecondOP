package Function

import (
	"database/sql"
	"fmt"
)

type MD5Struct struct {
	NTLM     string
	Password string
}

func Md5_query(hash_str string, DB sql.DB) MD5Struct {
	var user MD5Struct
	err := DB.QueryRow("SELECT * FROM md5 WHERE ntlm = ?", hash_str).Scan(&user.NTLM, &user.Password)
	if err != nil {
		fmt.Println("查询出错了")
	}
	return user
}
func MD5_insert(hash_str string, password string, DB sql.DB) {
	result, err := DB.Exec("insert into md5(ntlm,password) values(?,?)", hash_str, password)
	if err != nil {
		fmt.Println("新增数据错误", err)
		return
	}
	if result != nil {
		return

	}
}
