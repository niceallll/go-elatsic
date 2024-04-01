package logic

import (
	"webesapp/dao/mysql"
	"webesapp/models"
)

func CreateV6Qurey(p *models.EsSqlDetail) (err error) {
	err = mysql.CreateV6Qurey(p)
	if err != nil {
		return err
	}
	return
}
