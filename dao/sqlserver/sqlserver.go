package sqlserver

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	setting "webesapp/settings"
)

var db *sql.DB

func Init(cfg *setting.SQLserverConfig) (err error) {
	connstr := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s", cfg.Host, cfg.Port, cfg.DB, cfg.User, cfg.Password)
	db, err = sql.Open("mssql", connstr)
	if err != nil {
		return
	}
	return
}

func Close() {
	_ = db.Close()
}
