package adapter

import (
	"fmt"

	"github.com/go-rel/mysql"
	"github.com/go-rel/rel"
	mysqlDriver "github.com/go-sql-driver/mysql"
)

func GetRelAdapter() (rel.Adapter, error) {
	dsn := getMysqlConfig().FormatDSN()

	adapter, err := mysql.Open(dsn)
	if err != nil {
		return nil, err
	}

	return adapter, nil
}

func getMysqlConfig() *mysqlDriver.Config {

	host := "localhost"
	port := 3306
	user := "root"
	pass := "1234"
	dbName := "pocorm"

	config := mysqlDriver.NewConfig()
	config.Net = "tcp"
	config.Addr = fmt.Sprintf("%s:%d", host, port)
	config.DBName = dbName
	config.User = user
	config.Passwd = pass
	config.ParseTime = true

	return config
}
