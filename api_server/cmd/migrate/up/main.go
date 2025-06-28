package up

import (
	driver_conn "api/internal/io_infra/database/driver"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func SetUp() {

	db, _ := driver_conn.NewMySqlDriver()
	defer db.Close()
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

	fmt.Println(m)
	no, is_dirty, _ := m.Version()

	if is_dirty {
		m.Force(int(no))
	}

	if err = m.Up(); err != nil {
		fmt.Println("up err => ", err.Error())

		return
	}

	fmt.Println("テーブルのマイグレーションが完了しました。")
}
