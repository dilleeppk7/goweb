package acl

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Acl_user() {
	db_opts := map[string]string{
		"user":     "D",
		"password": "d123",
		"database": "mktplace",
		"host":     "localhost",
		"port":     "3306",
	}

	db, err := DbInit(db_opts)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	rows, err := db.Query("SELECT * FROM acl_user")
	if err != nil {
		fmt.Println(err)
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Print(columns[i], ": ", value)
			fmt.Print("\t")
		}
		fmt.Println("")
	}

}
