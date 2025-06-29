package up

import (
	driver_conn "api/internal/io_infra/database/driver"

	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var is_migrate_done = false

func SetUp() {

	db, _ := driver_conn.NewMySqlDriver()
	defer func() {
		recover()
		if !is_migrate_done {
			timer := time.NewTimer(15 * time.Second)
			<-timer.C
			SetUp()
		}

	}()
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:///usr/src/app/internal/io_infra/database/migration",
		"mysql",
		driver,
	)

	if err != nil {
		fmt.Println("migrateion err => ", err.Error())
		return
	}

	if m == nil {
		fmt.Println("m is nil")
		return
	}

	no, is_dirty, _ := m.Version()

	if is_dirty {
		m.Force(int(no))
	}

	if err = m.Up(); err != nil {
		fmt.Println("up err => ", err.Error())

		return
	}
	is_migrate_done = true

	fmt.Println("テーブルのマイグレーションが完了しました。")
}
