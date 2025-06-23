package driver

import (
	my_sql "api/internal/io_infra/config/my_sql_config"
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
)

//-------------------------------[driver interface]----------------------------//

// RepositoryはこのDriverインターフェースに依存

type Driver[DriverType sql.DB | any] interface {
	GetDriver() (*DriverType, error)
	Close()
	IsError() error
}

//-------------------------------[driver interface]----------------------------//

/*
	以下はdriver ・ database clientの実装
*/

//-----------------------------[mysql-sql.DB]------------------------------------//

type MySqlDriver[DriverType sql.DB | any] struct {
	db  *DriverType
	err error
}

func (msd *MySqlDriver[DriverType]) GetDriver() (*DriverType, error) {
	if msd.err != nil {
		return nil, msd.err
	}
	return msd.db, nil
}

func (msd *MySqlDriver[DriverType]) Close() {
	if msd.err != nil {
		return
	}

}

func MySqlUrl() string {
	// sql.Open("mysql", "mysql_user:todo_database_password@(localhost:3306)/todo_database")
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
		my_sql.DB_USER,
		my_sql.DB_PASSWORD,
		my_sql.DB_ADDR,
		my_sql.DB_PORT,
		my_sql.DB_NAME,
	)
}

func NewMySqlDriver() (*sql.DB, error) {
	url := MySqlUrl()
	new_db, err := sql.Open("mysql", url)
	new_db.SetMaxOpenConns(my_sql.DB_MAX_CONN)
	new_db.SetMaxIdleConns(my_sql.DB_MAX_CONN)
	if err != nil {
		ctx := context.Background()
		slog.Log(ctx, slog.LevelInfo, err.Error())
		return nil, err
	}
	return new_db, nil
}
