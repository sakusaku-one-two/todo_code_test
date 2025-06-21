package driver

import (
	"github.com/go-sql-driver/mysql"
)

// --------------------------[抽象的なコネクションハンドラーの形式　レポジトリはこの形式に依存]--------------------------//

type ConnectionHandler[DriverType any] interface {
	GetDriver() DriverType
	Close()
	Ping() bool
}

// --------------------------[mysql用のdb clientの作成]--------------------------//

type DriverTypeAsMySql = mysql.MySQLDriver

/*
	レポジトリが ConnectionHandlerを扱う際には　DriverTypeAsMySqlをジェネリックに差し込む
	サンプル：
		 ConnectionHandler[database_driver.DriverTypeAsMySql]

*/

type MySqlConnectionHandler struct {
	driver *mysql.MySQLDriver
}

func NewMySqlConnectionHandler() (ConnectionHandler[DriverTypeAsMySql], error) {

	mysql_config := mysql.Config{}

	driver, err := mysql.MySQLDriver(mysql_config)
	if err != nil {
		return nil, err
	}

	my_sql_handler := &MySqlConnectionHandler{
		driver: driver,
	}

	return my_sql_handler, nil
}

func (msh *MySqlConnectionHandler) GetDriver() *mysql.MySQLDriver {
	return msh.driver
}

func (msh *MySqlConnectionHandler) Close() {
	msh.driver.Close()
}

func (msh *MySqlConnectionHandler) Ping() bool {
	// 後程実装
	return true
}
