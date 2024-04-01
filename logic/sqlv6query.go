package logic

import (
	"webesapp/dao/mysql"
	"webesapp/dao/sqlserver"
	"webesapp/models"
)

func SqlV6Qurey() (p []*models.EsSqlDetail, err error) {
	return mysql.SqlV6Qurey()
}

func SqlV7Qurey() (p []*models.EsSqlDetail, err error) {
	return mysql.SqlV7Qurey()
}

func SSqlV6Qurey() (p []*models.EsSqlserverDetail, err error) {
	return sqlserver.SSqlV6Qurey()
}
