package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"webesapp/dao/es"
	"webesapp/models"
)

func CreateV6Qurey(p *models.EsSqlDetail) (err error) {
	sqlstr := `insert into crontab_v6(es_name,es_index,es_most_msg,es_most_not_msg,time,ding_talk,tlp,size,aliname,warn_up)values(?,?,?,?,?,?,?,?,?,?)`
	_, err = db.Exec(sqlstr, p.Esname, p.EsIndex, p.EsMostMsg, p.EsMostNotMsg, p.Time, p.DingTalk, p.Tlp, p.Size, p.AliName, p.Warn_up)

	if err != nil {
		zap.L().Error("create sqlv6qurey ", zap.Error(err))
	}
	zap.L().Debug("V6条件添加成功")
	return
}

func UpdateV6Qurey(p *models.EsSqlDetail) (err error) {
	sqlstr := `UPDATE   crontab_v6 set es_name=?, 
                        set es_index=?, 
                        set es_most_msg=?,
                        set es_most_not_msg=?,
                        set time=?,
                        set ding_talk=? ,
                        set tlp=?,
                        set size=?,
                        set aliname=?,
                        set warn_up=? where id=? `
	_, err = db.Exec(sqlstr, p.Esname, p.EsIndex, p.EsMostMsg, p.EsMostNotMsg, p.Time, p.DingTalk, p.Tlp, p.Size, p.AliName, p.Warn_up, p.ID)

	if err != nil {
		zap.L().Error("create sqlv6qurey ", zap.Error(err))
	}
	zap.L().Debug("V6条件添加成功")
	return
}

func SqlV6Qurey() (p []*models.EsSqlDetail, err error) {
	sqlstr := `select id,es_name,es_index,es_most_msg,es_most_not_msg,size, time,ding_talk,tlp,COALESCE(aliname, '') AS aliname,warn_up from crontab_v6  WHERE del = 1`
	if err := db.Select(&p, sqlstr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no crontab in db")
			err = nil
		}
	}

	for _, v := range p {
		if len(v.Esname) == 0 {
			continue
		}
		if v.Warn_up == 0 {
			continue
		}
		//if v.Del == 0 {
		//	continue
		//}
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
		v.Del = 1
		es.PostV6Qurey(m)
	}
	fmt.Println(&p)
	return
}
func SqlV7Qurey() (p []*models.EsSqlDetail, err error) {
	sqlstr := `select id,es_name,es_index,es_most_msg,size,es_most_not_msg,time,ding_talk,tlp,COALESCE(aliname, '') AS aliname,warn_up  from crontab_v7  WHERE del = 1`
	if err := db.Select(&p, sqlstr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no crontab in db")
			err = nil
		}
	}

	for _, v := range p {
		if len(v.Esname) == 0 {
			continue
		}
		if v.Warn_up == 0 {
			continue
		}
		if v.Del == 0 {
			continue
		}
		if len(v.AliName) == 0 {
			v.AliName = "空"
		}
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
		m.ID = v.ID
		v.Del = 1
		go es.PostV7Qurey(m)
		fmt.Println(m)
	}
	fmt.Println(&p)
	return
}

func CreateV7Qurey(p *models.EsSqlDetail) (err error) {
	sqlstr := `insert into crontab_v7(es_name,es_index,es_most_msg,es_most_not_msg,time,ding_talk,tlp,size,aliname,warn_up,del)values(?,?,?,?,?,?,?,?,?,?,?)`
	_, err = db.Exec(sqlstr, p.Esname, p.EsIndex, p.EsMostMsg, p.EsMostNotMsg, p.Time, p.DingTalk, p.Tlp, p.Size, p.AliName, p.Warn_up, p.Del)

	if err != nil {
		zap.L().Error("create sqlv7qurey ", zap.Error(err))
	}
	zap.L().Debug("V7条件添加成功")
	return
}

func CreateQurey(p *models.EsSqlDetail) error {
	sqlstr := `SELECT es_version FROM es_version WHERE es_name = ?`
	var esVersion string
	err := db.QueryRow(sqlstr, p.Esname).Scan(&esVersion)
	if err != nil {
		zap.L().Error("查询es_version失败", zap.Error(err))
		return err
	}

	switch esVersion {
	case "v6":
		sqlstr := `INSERT INTO crontab_v6(es_name, es_index, es_most_msg, es_most_not_msg, time, ding_talk, tlp, size, aliname, warn_up,del) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		_, err = db.Exec(sqlstr, p.Esname, p.EsIndex, p.EsMostMsg, p.EsMostNotMsg, p.Time, p.DingTalk, p.Tlp, p.Size, p.AliName, p.Warn_up, p.Del)
		if err != nil {
			zap.L().Error("插入crontab_v6数据失败", zap.Error(err))
			return err
		}
		fmt.Println(p.Esname, p.EsIndex, p.EsMostMsg, p.EsMostNotMsg, p.Time, p.DingTalk, p.Tlp, &p.Size, p.AliName, &p.Warn_up)
		zap.L().Debug("V6条件添加成功")
	case "v7":
		sqlstr := `INSERT INTO crontab_v7(es_name, es_index, es_most_msg, es_most_not_msg, time, ding_talk, tlp, size, aliname, warn_up,del) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		_, err = db.Exec(sqlstr, p.Esname, p.EsIndex, p.EsMostMsg, p.EsMostNotMsg, p.Time, p.DingTalk, p.Tlp, p.Size, p.AliName, p.Warn_up, p.Del)
		if err != nil {
			zap.L().Error("插入crontab_v7数据失败", zap.Error(err))
			return err
		}
		zap.L().Debug("V7条件添加成功")
	default:
		zap.L().Error("未知的es_version:", zap.String("es_version", esVersion))
		return fmt.Errorf("未知的es_version: %s", esVersion)
	}

	return nil
}

func UpdateQurey(p *models.EsSqlDetail) error {
	sqlstr := `SELECT es_version FROM es_version WHERE es_name = ?`
	var esVersion string
	err := db.QueryRow(sqlstr, p.Esname).Scan(&esVersion)
	if err != nil {
		zap.L().Error("修改es_version失败", zap.Error(err))
		return err
	}

	switch esVersion {
	case "v6":
		sqlstr := `UPDATE   crontab_v6 set es_name=?, 
                        es_index=?, 
                        es_most_msg=?,
                        es_most_not_msg=?,
                        time=?,
                        ding_talk=? ,
                        tlp=?,
                        size=?,
                        aliname=?,
                        warn_up=? where id=? `
		_, err = db.Exec(sqlstr, p.Esname, p.EsIndex, p.EsMostMsg, p.EsMostNotMsg, p.Time, p.DingTalk, p.Tlp, p.Size, p.AliName, p.Warn_up, p.ID)
		if err != nil {
			zap.L().Error("修改crontab_v6数据失败", zap.Error(err))
			return err
		}
		fmt.Println(p.Esname, p.EsIndex, p.EsMostMsg, p.EsMostNotMsg, p.Time, p.DingTalk, p.Tlp, &p.Size, p.AliName, &p.Warn_up)
		zap.L().Debug("V6条件添加成功")
	case "v7":
		sqlstr := `UPDATE   crontab_v7 set es_name=?, 
                        es_index=?, 
                        es_most_msg=?,
                        es_most_not_msg=?,
                        time=?,
                        ding_talk=? ,
                        tlp=?,
                        size=?,
                        aliname=?,
                        warn_up=? where id=? `
		_, err = db.Exec(sqlstr, p.Esname, p.EsIndex, p.EsMostMsg, p.EsMostNotMsg, p.Time, p.DingTalk, p.Tlp, p.Size, p.AliName, p.Warn_up, p.ID)
		if err != nil {
			zap.L().Error("修改crontab_v7数据失败", zap.Error(err))
			return err
		}
		zap.L().Debug("V7条件修改成功")
	default:
		zap.L().Error("未知的es_version:", zap.String("es_version", esVersion))
		return fmt.Errorf("未知的es_version: %s", esVersion)
	}

	return nil
}

func DeleteQurey(p *models.EsSqlDetail) (err error) {
	sqlstr := `SELECT es_version from es_version WHERE es_name = ?`
	var esVersion string
	err = db.QueryRow(sqlstr, p.Esname).Scan(&esVersion)
	if err != nil {
		zap.L().Error("查询es_version失败", zap.Error(err))
		return
	}

	switch esVersion {
	case "v6":
		sqlstrde := `UPDATE   crontab_v6 set del = 0  WHERE id = ?`
		_, err = db.Exec(sqlstrde, p.ID)
		if err != nil {
			zap.L().Error("删除v6版本记录", zap.Error(err))
			return err
		}
		zap.L().Debug("v6记录删除成功")
	case "v7":
		sqlstrde := `UPDATE  crontab_v7 set del = 0  WHERE id = ?`
		_, err = db.Exec(sqlstrde, p.ID)
		if err != nil {
			zap.L().Error("删除v7版本记录", zap.Error(err))
			return err
		}
		zap.L().Debug("v7记录删除成功")
	default:
		zap.L().Error("未知的es_version:", zap.String("es_version", esVersion))
		return fmt.Errorf("未知的es_version: %s", esVersion)
	}
	return nil

}
