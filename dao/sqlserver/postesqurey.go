package sqlserver

import (
	"encoding/json"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"go.uber.org/zap"
	"webesapp/dao/es"
	"webesapp/models"
)

func CreatesqlV6Qurey(p *models.EsSqlserverDetail) (err error) {
	sqlstr := `insert into crontab(es_name,es_index,es_most_msg,es_most_not_msg,time,ding_talk,tlp,size)values(?,?,?,?,?,?,?,?)`
	_, err = db.Exec(sqlstr, p.Esname, p.EsIndex, p.EsMostMsg, p.EsMostNotMsg, p.Time, p.DingTalk, p.Tlp, p.Size)
	zap.L().Error("create create sqlserver qurey ", zap.Error(err))
	return
}

func SSqlV6Qurey() (p []*models.EsSqlserverDetail, err error) {

	sqlStr := "SELECT id,size, es_name,es_index,tlp,time,ding_talk,es_most_msg,es_most_not_msg ,aliname FROM crontab"
	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	fmt.Println(rows)
	defer rows.Close()

	//var esDetail models.EsSqlserverDetail
	//err = rows.Scan(&esDetail.ID, &esDetail.Esname, &esDetail.EsIndex, &esDetail.EsMostNotMsg, &esDetail.EsMostMsg, &esDetail.Time, &esDetail.Size, &esDetail.DingTalk, &esDetail.Tlp)
	//if err != nil {
	//	return
	//}

	for rows.Next() {

		var v models.EsSqlserverDetail
		err = rows.Scan(
			&v.ID,
			&v.Size,
			&v.Esname,
			&v.EsIndex,
			&v.Tlp,
			&v.Time,
			&v.DingTalk,
			&v.EsMostMsg,
			&v.EsMostNotMsg,
			&v.AliName,
		)

		m := new(models.CatNginxStr)
		json.Unmarshal([]byte(v.EsMostNotMsg), &m.EsMustNot)
		json.Unmarshal([]byte(v.EsMostMsg), &m.EsMust)
		json.Unmarshal([]byte(v.EsIndex), &m.EsIndex)
		json.Unmarshal([]byte(v.DingTalk), &m.Dingtalk)
		json.Unmarshal([]byte(v.Tlp), &m.Tlp)
		m.Size = v.Size
		m.Time = v.Time
		m.EsName = v.Esname
		m.ALiName = v.AliName

		go es.PostV6Qurey(m)
		fmt.Println(m)

		p = append(p, &v)
	}

	fmt.Println(&p)
	return p, nil
}
