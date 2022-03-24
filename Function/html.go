package Function

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Line struct {
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
	SP int     `json:"sp"`
	AG int     `json:"ag"`
	TM int     `json:"tm"`
}

func Linequery(DB *sql.DB, w http.ResponseWriter, name string, starttime string, endtime string) (string, error) {

	var data []Line

	rows, err := DB.Query("SELECT X,Y,SP,AG,TM FROM record WHERE line = ? and TM <? and TM > ? ORDER BY TM", name, endtime, starttime)
	if err != nil {
		fmt.Println("select db failed,err:", err)

		return "fail", errors.New("failed")
	}
	for rows.Next() {
		var user Line
		rows.Scan(&user.X, &user.Y, &user.SP, &user.AG, &user.TM)
		data = append(data, user)

	}

	t, _ := template.New("line.html").ParseFiles("html/line.html")
	for i := 1; i < len(data); i++ {
		data[i].TM = data[i].TM - data[0].TM
	}

	t.Execute(w, data)
	rows.Close()

	return "success", nil

}

func Lineinsert(DB *sql.DB, X string, Y string, AG string, name string) (string, error) {

	_, err := DB.Exec("insert into androids.record(X,Y,SP,AG,TM,line) values(?,?,10,?,?,?)", X, Y, AG, time.Now().Unix(), name)
	if err != nil {
		fmt.Println("新增数据错误", err)

		return "inserterror", errors.New("insert error")
	}

	return "success", nil
}

func Linelist(DB *sql.DB) ([]string, error) {

	var Horselist []string

	rows, _ := DB.Query("SELECT line from record GROUP BY line")

	for rows.Next() {
		var user string
		rows.Scan(&user)
		Horselist = append(Horselist, user)
	}
	rows.Close()

	return Horselist, nil
}
