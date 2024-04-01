package logic

import (
	"webesapp/dao/mysql"
	"webesapp/models"
)

func UpdateV6Qurey(p *models.EsSqlDetail) (err error) {
	err = mysql.UpdateV6Qurey(p)
	if err != nil {
		return err
	}
	return
}
func UpdateQurey(p *models.EsSqlDetail) (err error) {
	err = mysql.UpdateQurey(p)
	if err != nil {
		return err
	}
	return
}
