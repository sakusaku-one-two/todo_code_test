package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// MYSQL_ROOT_PASSWORD=root_password
// MYSQL_DATABASE=todo_database
// MYSQL_USER=mysql_user
// MYSQL_PASSWORD=todo_database_password

func NewD() (*sql.DB, error) {
	new_db, err := sql.Open("mysql", "mysql_user:todo_database_password@localhost:3306/todo_database")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return new_db, nil
}
