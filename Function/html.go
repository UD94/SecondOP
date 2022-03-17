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

func Lineinsert(X float64, Y float64, AG int, name string) (string, error) {
	DB := new(sql.DB)
	defer DB.Close()
	err := Common.InitDB(DB, "Lines")
	if err == nil {
		_, err := DB.Exec("insert into lines(X,Y,SP,AG,TM,line) values(?,?,10,?,?,?)", X, Y, AG, time.Now().Unix(), name)
		if err != nil {
			fmt.Println("新增数据错误", err)
			return "inserterror", errors.New("insert error")
		}
	} else {
		return "nodatabase", errors.New("no database")
	}
	return "success", nil
}
