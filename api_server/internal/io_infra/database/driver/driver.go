package driver

import (
	my_sql "api/internal/io_infra/config/my_sql_config"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func MySqlUrl() string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true&multiStatements=true",
		my_sql.DB_USER,
		my_sql.DB_PASSWORD,
		my_sql.DB_ADDR,
		my_sql.DB_PORT,
		my_sql.DB_NAME,
	)
}

func NewMySqlDriver() (*sql.DB, error) {
	url := MySqlUrl()
	defer recover()
	new_db, err := sql.Open("mysql", url)
	// new_db.SetMaxOpenConns(my_sql.DB_MAX_CONN)
	// new_db.SetMaxIdleConns(my_sql.DB_MAX_CONN)

	if err != nil {
		// ctx := context.Background()
		// slog.Log(ctx, slog.LevelInfo, err.Error())
		time_chan := time.NewTimer(10 * time.Second).C
		<-time_chan
		return NewMySqlDriver()
	}
	return new_db, nil
}
