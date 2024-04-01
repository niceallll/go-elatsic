package logic

import (
	"webesapp/dao/mysql"
	"webesapp/models"
)

func CreateV7Qurey(p *models.EsSqlDetail) (err error) {
	err = mysql.CreateV7Qurey(p)
	if err != nil {
		return err
	}
	return
}
