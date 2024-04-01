package es

import (
	"context"
	"fmt"
	"github.com/blinkbean/dingtalk"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
	"webesapp/models/aliyunopsfile"

	//"github.com/olivere/elastic"
	"github.com/olivere/elastic/v7"
	"reflect"
	"webesapp/models"
	"webesapp/models/cos"
)

var resv7 *elastic.SearchResult

type V7Catnginxidx struct {
	Status      int64  `json:"status"`      // 状态码
	Request_uri string `json:"request_uri"` // 请求地址
	Domain      string `json:"domain"`      // 域名
	Stack_trace string `json:"stack_trace"` //
	Msg         string `json:"msg"`         // 内容
	Level       string `json:"level"`       // 级别
	App         string `json:"app"`         // 程序
}

func (v V7Catnginxidx) catalinav7(resv7 *elastic.SearchResult, indx, esname string, size int64, ding, tlp []string, aliname string) (ts []models.CatNginxidx) {
	//var dingToken = []string{""}
	if size == 0 {
		size = 4
	}
	if ding == nil && tlp == nil {
		ding = []string{"xxxxx"}
		tlp = []string{"xxxx"}
	}
	var dingTalkCli = dingtalk.InitDingTalk(ding, ".")
	//mobilbs := []string{"17612735074"}
	var aa string
	bb := []string{}
	cc := esname + "" + "\n" + "**索引" + indx + "\n"
	if resv7.TotalHits() > size {
		for _, item := range resv7.Each(reflect.TypeOf(v)) {
			if t, ok := item.(V7Catnginxidx); ok {
				aa = aa + "- <font color=#00ff00 size=6>catalinamsg</font>" + "   " + t.Msg + "\n"
			}
			if strings.Count(aa, "catalinamsg") > 4 {
				goto Loop
			}
		}
	Loop:
		bb = append(bb, cc+aa)
		//dingTalkCli.SendMarkDownMessageBySlice("ERROR MSG", bb, dingtalk.WithAtMobiles(tlp))
		//fmt.Println(bb)
		fmt.Println(resv)
		//os.Mkdir("/Users/wenchao.ma/go/data/logs", os.ModePerm)
		var urls *url.URL
		rand.Seed(time.Now().UnixNano())
		//currentTime := time.Now().Format("2006-01-0215:04:05.0000000")
		randomInt := rand.Intn(1000000)
		currentTime := strconv.Itoa(randomInt)
		filename := esname + currentTime + ".html"
		filePtr, err := os.Create(filename)
		if err != nil {
			fmt.Println("文件创建失败", err.Error())
			return
		}
		defer filePtr.Close()
		filePtr.WriteString("<!DOCTYPE html><html lang=\"en\"><head>    <meta charset=\"UTF-8\">    <title>" + esname + "</title></head><body>")
		var typ []models.CatNginxidx
		if resv7.TotalHits() > size {
			var typ1 models.CatNginxidx
			filePtr.WriteString("<table border='1'>")
			for _, item := range resv7.Each(reflect.TypeOf(typ1)) {
				if t, ok := item.(models.CatNginxidx); ok {
					//typ = []models.CatNginxidx{}
					typ1.Msg = t.Msg
					typ1.Stack_trace = t.Stack_trace
					typ1.Level = t.Level
					typ1.App = t.App
					typ1.Params = t.Params
					typ = append(typ, typ1)
					filePtr.WriteString("<tr>")
					filePtr.WriteString("<td>" + "Msg" + "</td>")
					filePtr.WriteString("<td>" + t.Msg + "</td> \n")
					filePtr.WriteString("</tr>")

					filePtr.WriteString("<tr>")
					filePtr.WriteString("<td>" + "Stack_trace" + "</td>")
					filePtr.WriteString("<td>" + t.Stack_trace + "</td> \n")
					filePtr.WriteString("</tr>")

					filePtr.WriteString("<tr>")
					filePtr.WriteString("<td>" + "Params" + "</td>")
					filePtr.WriteString("<td>" + t.Params + "</td> \n")
					filePtr.WriteString("</tr>")

					filePtr.WriteString("<tr>")
					filePtr.WriteString("<td>" + "APP" + "</td>")
					filePtr.WriteString("<td>" + t.App + "</td> \n")
					filePtr.WriteString("</tr>")

					filePtr.WriteString("<tr>")
					filePtr.WriteString("<td>" + "Level" + "</td>")
					filePtr.WriteString("<td>" + t.Level + "</td> \n")
					filePtr.WriteString("</tr>")
				}
			}
			filePtr.WriteString("</table>")
			urls = cos.Upload(filename)
			s := urls.String()
			sa := "详细信息: [日志链接]" + "(" + s + ")"
			bb = append(bb, sa)
			//dingTalkCli.SendMarkDownMessage("URL", sa, dingtalk.WithAtMobiles(tlp))
			dingTalkCli.SendMarkDownMessageBySlice("ERROR MSG", bb, dingtalk.WithAtMobiles(tlp))
			aliyunopsfile.Openfile(sa, aliname)
			bb = []string{}
		}
		return typ
	}
	return
}
func (v V7Catnginxidx) nginxv7(resv7 *elastic.SearchResult, indx, esname string, size int64, ding, tlp []string, aliname string) (ts []models.CatNginxidx) {

	if size == 0 {
		size = 10
	}
	if ding == nil && tlp == nil {
		ding = []string{"xxxxxx"}
		tlp = []string{"xxxx"}
	}
	var dingTalkCli = dingtalk.InitDingTalk(ding, ".")
	var aa string
	bb := []string{}
	cc := esname + "\n" + "*索引名称： " + indx + "\n"
	if resv7.TotalHits() > size {
		for _, item := range resv7.Each(reflect.TypeOf(v)) {
			if t, ok := item.(V7Catnginxidx); ok {
				aa = aa + "- <font color=#00ff00 size=6>nginxmsg</font>" + "   " + t.Request_uri + "\n"
			}
			if strings.Count(aa, "nginxmsg") > 4 {
				goto Loop
			}
		}
	Loop:
		bb = append(bb, cc+aa)
		fmt.Println(resv)
		rand.Seed(time.Now().UnixNano())
		//currentTime := time.Now().Format("2006-01-0215:04:05.0000000")
		randomInt := rand.Intn(1000000)
		currentTime := strconv.Itoa(randomInt)
		filename := esname + currentTime + ".html"
		filePtr, err := os.Create(filename)
		if err != nil {
			fmt.Println("文件创建失败", err.Error())
			return
		}
		defer filePtr.Close()
		filePtr.WriteString("<!DOCTYPE html><html lang=\"en\"><head>    <meta charset=\"UTF-8\">    <title>" + esname + "</title></head><body>")
		var typ []models.CatNginxidx
		if resv7.TotalHits() > size {
			var typ1 models.CatNginxidx
			filePtr.WriteString("<table border='1'>")
			for _, item := range resv7.Each(reflect.TypeOf(typ1)) {
				if t, ok := item.(models.CatNginxidx); ok {
					typ1.App = t.App
					//typ = []models.CatNginxidx{}
					typ1.Request_uri = t.Request_uri
					typ1.Status = t.Status
					typ1.Domain = t.Domain
					typ1.App = t.App
					str2 := fmt.Sprintf("%d", t.Status)
					filePtr.WriteString("<tr>")
					filePtr.WriteString("<td>" + "Request_uri" + "</td>")
					filePtr.WriteString("<td>" + t.Request_uri + "</td> \n")
					filePtr.WriteString("</tr>")
					filePtr.WriteString("<tr>")
					filePtr.WriteString("<td>" + "Domain" + "</td>")
					filePtr.WriteString("<td>" + t.Domain + "</td> \n")
					filePtr.WriteString("</tr>")
					filePtr.WriteString("<tr>")
					filePtr.WriteString("<td>" + "Status" + "</td>")
					filePtr.WriteString("<td>" + str2 + "</td> \n")
					filePtr.WriteString("</tr>")
					typ = append(typ, typ1)
				}
			}

			filePtr.WriteString("</table>")
			urls := cos.Upload(filename)
			s := urls.String()
			sa := "详细信息: [日志链接]" + "(" + s + ")"
			bb = append(bb, sa)
			dingTalkCli.SendMarkDownMessageBySlice("NGINX", bb, dingtalk.WithAtMobiles(tlp))
			aliyunopsfile.Openfile(sa, aliname)
			bb = []string{}
		}
		return typ
	}
	return
}

//处理catalina请求
func PostV7Qurey(es *models.CatNginxStr) (ts *[]models.CatNginxidx, err error) {
	var ding []string
	for _, v := range es.Dingtalk {
		ding = append(ding, v)
	}

	var tlp []string
	for _, v := range es.Tlp {
		tlp = append(tlp, v)
	}
	aliname := es.ALiName
	var v = &V7Catnginxidx{}
	var typ []models.CatNginxidx
	m, _ := time.ParseDuration(es.Time)
	//fmt.Println(m)
	newTime := time.Now()         //.Format("2006-01-02T15:04:05")
	downtime := time.Now().Add(m) //.Format("2006-01-02T15:04:05")
	boolcaQuery := elastic.NewBoolQuery()
	boolcaQuery.Must(elastic.NewRangeQuery("@timestamp").Gte(downtime).Lte(newTime))
	for k, v := range es.EsMust {
		for _, i := range v {
			boolcaQuery.Must(elastic.NewMatchQuery(k, i))
		}
	}
	for k, v := range es.EsMustNot {
		for _, i := range v {
			boolcaQuery.MustNot(elastic.NewMatchPhraseQuery(k, i))
		}
	}
	for _, indx := range es.EsIndex {
		b := strings.Contains(indx, "nginx")
		if b == true {
			resv7, err = esclientV7[es.EsName].
				Search().
				Index(indx).
				Query(boolcaQuery).
				Size(50).
				Pretty(true).
				Do(context.Background())
			fmt.Println(boolcaQuery)
			typ = v.nginxv7(resv7, indx, es.EsName, es.Size, ding, tlp, aliname)
		} else {
			resv7, err = esclientV7[es.EsName].
				Search().
				Index(indx).
				Query(boolcaQuery).
				Size(100).
				Pretty(true).
				Do(context.Background())
			typ = v.catalinav7(resv7, indx, es.EsName, es.Size, ding, tlp, aliname)

		}
	}
	return &typ, err
}
