package cinit

import (
	"strconv"

	"github.com/jmoiron/sqlx"
)

var Mysql *sqlx.DB

func mysqlInit() {
	var err error
	dataSourceName := Config.Mysql.User + ":" + Config.Mysql.Password + "@tcp(" + Config.Mysql.Addr + ":" + strconv.Itoa(Config.Mysql.Port) +
		")/" + Config.Mysql.DbName + "?parseTime=true&loc=Local"
	Mysql, err = sqlx.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	Mysql.SetMaxIdleConns(Config.Mysql.IDleConn)
	Mysql.SetMaxOpenConns(Config.Mysql.MaxConn)
	err = Mysql.Ping()
	if err != nil {
		panic(err)
	}
}

func mysqlClose() {
	Mysql.Close()
}
