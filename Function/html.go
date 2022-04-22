package Function

import (
	"bytes"
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

type Androidinfo struct {
	Line    string
	Job     string
	Message string
	Reverso string
}

func Linequery(DB *sql.DB, w http.ResponseWriter, name string, starttime string, endtime string) (string, error) {

	var data []Line

	rows, err := DB.Query("SELECT X,Y,SP,AG,TM FROM record WHERE line = ? and TM <? and TM > ? ORDER BY TM ", name, endtime, starttime)
	if err != nil {
		fmt.Println("select db failed,err:", err)

		return "fail", errors.New("failed")
	}
	for rows.Next() {
		var user Line
		rows.Scan(&user.X, &user.Y, &user.SP, &user.AG, &user.TM)
		data = append(data, user)

	}
	var buf bytes.Buffer

	t, _ := template.New("line.html").ParseFiles("html/line.html")
	for i := 1; i < len(data); i++ {
		data[i].TM = data[i].TM - data[0].TM
	}

	t.Execute(&buf, data)
	rows.Close()

	return buf.String(), nil

}

func Androidjob(DB *sql.DB, target string, value string) (string, error) {
	_, err := DB.Exec("UPDATE androids.target SET line= ?, job= ?  WHERE (line= ? )  LIMIT 1;	", target, value, target)
	if err != nil {
		fmt.Println("新增数据错误", err)

		return "updateerror", errors.New("update error")
	}
	return "success", nil
}

func Androidjobqeury(DB *sql.DB, target string) (string, error) {
	var user Androidinfo
	err := DB.QueryRow("select * from target where line = ? ", target).Scan(&user.Line, &user.Job, &user.Message, &user.Reverso)
	if err != nil {
		fmt.Println("查询出错了", err)

		return "nopass", errors.New("no pass")
	}

	return user.Job, nil
}

func Lineinsert(DB *sql.DB, X string, Y string, AG string, name string, Useragent string) (string, error) {

	_, err := DB.Exec("insert into androids.record(X,Y,SP,AG,TM,line,UA) values(?,?,10,?,?,?,?)", X, Y, AG, time.Now().Unix(), name, Useragent)
	if err != nil {
		fmt.Println("新增数据错误", err)

		return "inserterror", errors.New("insert error")
	}
	lists, _ := Linelist(DB)
	if Common.In(name, lists) {
		DB.Exec("insert into androids.target(line,job,message,Reverso) values(?,?,?,?)", name, "0", "empty", "empty")

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
