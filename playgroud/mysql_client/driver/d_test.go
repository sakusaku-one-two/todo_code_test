package driver

import (
	"fmt"
	"testing"
)

func Test_d(t *testing.T) {
	db, err := NewD()

	if err = db.Ping(); err != nil {
		// do not connect db
		fmt.Println(err)
	}

}
