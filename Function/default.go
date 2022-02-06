package Function

import "database/sql"

func Save() {

	DB := new(sql.DB)
	defer DB.Close()
}

func Read() {

	DB := new(sql.DB)
	defer DB.Close()
}
