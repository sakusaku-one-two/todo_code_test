package main

import (
	"fmt"
	"sample/driver"
)

func main() {
	db, err := driver.NewD()
	if err != nil {
		fmt.Println(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err)
	}
}
