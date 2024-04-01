package logic

import (
	"webesapp/dao/mysql"
	"webesapp/models"
)

func CreateQurey(p *models.EsSqlDetail) (err error) {
	err = mysql.CreateQurey(p)
	if err != nil {
		return err
	}
	return
}
