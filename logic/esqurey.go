package logic

import (
	"webesapp/dao/es"
	"webesapp/models"
)

//func GetV6nginxQurey(p string) (*[]models.CatNginxidx, error) {
//	loges, err := es.GetV6NginxQurey(p)
//	fmt.Println("logic es", loges, err)
//	return es.GetV6NginxQurey(p)
//}

func PostV6Qurey(ess *models.CatNginxStr) (ts *[]models.CatNginxidx, err error) {
	return es.PostV6Qurey(ess)
}

func PostV7Qurey(ess *models.CatNginxStr) (ts *[]models.CatNginxidx, err error) {
	return es.PostV7Qurey(ess)
}

//func PostV7NginxQurey(esname, estype, esmsg, esindex, t string) (*[]models.CatNginxidx, error) {
//	return es.PostV6NginxQurey(esname, estype, esmsg, esindex, t)
//}
