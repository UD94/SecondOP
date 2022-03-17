package Function

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/UD94/SecondOP/Common"
)

func RenderHTML(w http.ResponseWriter, file string, data interface{}) {
	// 获取页面内容
	t, _ := template.New(file).ParseFiles("html/" + file)

	// 将页面渲染后反馈给客户端
	t.Execute(w, data)
}

func Linequery(hash_str string) (string, error) {
	var user MD5Struct

	var DB *sql.DB
	DB, status := Common.InitDB("ntlm")
	if status == nil {
		err := DB.QueryRow("SELECT Password FROM ntlm WHERE line = ?", hash_str).Scan(&user.Hash, &user.Password)
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

func Lineinsert(X string, Y string, AG string, name string) (string, error) {
	var DB *sql.DB
	DB, err := Common.InitDB("Lines")
	if err == nil {
		_, err := DB.Exec("insert into Lines.lines(X,Y,SP,AG,TM,line) values(?,?,10,?,?,?)", X, Y, AG, time.Now().Unix(), name)
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
