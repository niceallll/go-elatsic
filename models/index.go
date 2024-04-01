package models

import (
	"bytes"
	"database/sql/driver"
)

type JSON []byte

func (j JSON) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {

	}
	*j = append((*j)[0:0], s...)
	return nil
}
func (m JSON) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}
func (m *JSON) UnmarshalJSON(data []byte) error {
	if m == nil {

	}
	*m = append((*m)[0:0], data...)
	return nil
}
func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}
func (j JSON) Equals(j1 JSON) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}

type Catalinaidx struct {
	Msg   string `json:"msg"`   // 内容
	Level string `json:"level"` // 级别
	App   string `json:"app"`   // 程序
}
type CatNginxStr struct {
	ID        int                 `json:"id"`
	EsName    string              `json:"esname"` // es数据库名称
	Time      string              `json:"time"`   // 查询时间
	Size      int64               `json:"size"`
	EsIndex   map[string]string   `json:"esindex"`     // 索引名称
	EsMust    map[string][]string `json:"es_must"`     //查询内容
	EsMustNot map[string][]string `json:"es_must_not"` //过滤内容
	Dingtalk  map[string]string   `json:"dingtalk"`
	Tlp       map[string]string   `json:"tlp"`
	ALiName   string              `json:"aliname"` //阿里云运维事件中心
	Warn      int                 `json:"warn"`
}
type CatNginxidx struct {
	Status      int    `json:"status"`      // 状态码
	Request_uri string `json:"request_uri"` // 请求地址
	Domain      string `json:"domain"`      // 域名
	Stack_trace string `json:"stack_trace"` //
	Msg         string `json:"msg"`         // 内容
	Level       string `json:"level"`       // 级别
	App         string `json:"app"`         // 程序
	Params      string `json:"params"`      // 携带内容
}

type NginxQurey struct {
	Status     int64  `json:"status" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type EsSqlDetail struct {
	ID           int    `json:"id" db:"id"`
	Esname       string `json:"es_name" db:"es_name"`
	EsIndex      JSON   `json:"es_index" db:"es_index"`
	EsMostMsg    JSON   `json:"es_most_msg" db:"es_most_msg"`
	EsMostNotMsg JSON   `json:"es_most_not_msg" db:"es_most_not_msg"`
	Time         string `json:"time" db:"time"`
	Size         int64  `json:"size" db:"size"`
	DingTalk     JSON   `json:"ding_talk" db:"ding_talk"`
	Tlp          JSON   `json:"tlp" db:"tlp"`
	AliName      string `json:"aliname" db:"aliname"`
	Warn_up      int    `json:"warn_up"db:"warn_up"`
	Del          int    `json:"del"db:"del"`
}

type EsSqlserverDetail struct {
	ID           int    `json:"id" db:"id"`
	Esname       string `json:"es_name" db:"es_name"`
	EsIndex      string `json:"es_index" db:"es_index"`
	EsMostMsg    string `json:"es_most_msg" db:"es_most_msg"`
	EsMostNotMsg string `json:"es_most_not_msg" db:"es_most_not_msg"`
	Time         string `json:"time" db:"time"`
	Size         int64  `json:"size" db:"size"`
	DingTalk     string `json:"ding_talk" db:"ding_talk"`
	Tlp          string `json:"tlp" db:"tlp"`
	AliName      string `json:"aliname" db:"aliname"`
}
