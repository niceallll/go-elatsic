package logic

import (
	"webesapp/dao/mysql"
	"webesapp/models"
)

func DeleteQurey(p *models.EsSqlDetail) (err error) {
	err = mysql.DeleteQurey(p)
	if err != nil {
		return err
	}
	return
}
