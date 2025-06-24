package driver

import (
	"fmt"
	"testing"
)

func Test_d(t *testing.T) {
	db, err := NewD()
	// defer func() {
	// 	err := recover()
	// 	fmt.Println(err)

	// }()

	if err = db.Ping(); err != nil {
		// do not connect db
		fmt.Println(err)
	} else {
		fmt.Println(db, "connection ok")
	}

	result, err := db.Exec("insert into todo(title,description) values('new_data','new_description');")

	rows, err := result.LastInsertId()
	fmt.Println(rows, err)
	fmt.Println("!!")
}
