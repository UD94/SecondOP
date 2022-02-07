package Function

import (
	"database/sql"

	"github.com/UD94/SecondOP/Common"
)

func Nmap() {

}

func Fscan() {

}

func Mimikatz(MimikatzString []string) {
	DB := new(sql.DB)
	defer DB.Close()
	first := Common.InitDB(DB, "ntlm")
	if first == nil {

	}
}

func AD() {

}
