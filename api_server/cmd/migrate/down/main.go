package main

import (
	driver_conn "api/internal/io_infra/database/driver"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, _ := driver_conn.NewMySqlDriver()
	defer db.Close()
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/io_infra/database/migration",
		"mysql",
		driver,
	)

	if err != nil {
		fmt.Println("migrateion err => ", err.Error())
		return
	}

	no, is_dirty, _ := m.Version()

	if is_dirty {
		m.Force(int(no))
	}

	if err = m.Down(); err != nil {
		fmt.Println("down err => ", err.Error())

		return
	}

	fmt.Println("MIGRATE DONE")
}
