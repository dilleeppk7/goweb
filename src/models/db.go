package acl

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func DbInit(options map[string]string) (sql.DB, error) {
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", options["user"], options["password"], options["host"], options["port"], options["database"])
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	} else {
		fmt.Println("DB connection established with ", connectionString)
	}
	defer db.Close()
	maxCon, ok := options["max"]
	if ok && maxCon != "" {
		maxcon, _ := strconv.ParseInt(maxCon, 10, 16)
		fmt.Println("using max connections ", maxcon)
		db.SetMaxOpenConns(int(maxcon))
	} else {
		fmt.Println("using default max connections ", 10)
		db.SetMaxOpenConns(10)
	}
	minCon, ok := options["min"]
	if ok && minCon != "" {
		mincon, _ := strconv.ParseInt(minCon, 10, 16)
		fmt.Println("using min connections ", mincon)
		db.SetMaxIdleConns(int(mincon))
	} else {

		fmt.Println("using default min connections ", 2)
		db.SetMaxIdleConns(2)
	}

	return *db, err
}
