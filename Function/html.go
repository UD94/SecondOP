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

type Line struct {
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
	SP int     `json:"sp"`
	AG int     `json:"ag"`
	TM int     `json:"tm"`
}

func Linequery2(w http.ResponseWriter, name string, starttime string, endtime string) (string, error) {

	var DB *sql.DB
	var data []Line
	DB, err := Common.InitDB("Lines")
	if err != nil {
		fmt.Println("connect to mysql failed,", err)
		return "fail", errors.New("failed")

	}
	rows, err := DB.Query("SELECT X,Y,SP,AG,TM FROM record WHERE line = ? and TM <? and TM > ? ORDER BY TM", name, endtime, starttime)
	if err != nil {
		fmt.Println("select db failed,err:", err)
		return "fail", errors.New("failed")
	}
	for rows.Next() {
		var user Line
		rows.Scan(&user.X, &user.Y, &user.SP, &user.AG, &user.TM)
		fmt.Println(user.X)
		data = append(data, user)

	}

	t, _ := template.New("line.html").ParseFiles("html/line.html")
	for i := 1; i < len(data); i++ {
		data[i].TM = data[i].TM - data[0].TM
	}
	fmt.Println(data)

	t.Execute(w, data)

	return "success", nil

}

func Lineinsert(X string, Y string, AG string, name string) (string, error) {
	var DB *sql.DB
	DB, err := Common.InitDB("Lines")
	if err == nil {
		_, err := DB.Exec("insert into Lines.record(X,Y,SP,AG,TM,line) values(?,?,10,?,?,?)", X, Y, AG, time.Now().Unix(), name)
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
