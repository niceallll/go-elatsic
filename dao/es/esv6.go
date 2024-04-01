package es

import (
	"fmt"
	//"fmt"
	"github.com/olivere/elastic"
	"go.uber.org/zap"
	setting "webesapp/settings"
)

var esclientV6 map[string]*elastic.Client

func InitV6(cfg *setting.ESconfig) (err error) {
	esclientV6 = make(map[string]*elastic.Client)
	for k, v := range cfg.Es6list {
		fmt.Println(v.IpHOST)
		fmt.Println(k)
		client, err := elastic.NewClient(elastic.SetURL(v.IpHOST), elastic.SetBasicAuth(v.User, v.Password), elastic.SetSniff(false))
		if err != nil {
			zap.L().Error("ES init fiald", zap.Error(err))
		}
		esversion, err := client.ElasticsearchVersion(v.IpHOST)
		if err != nil {
			zap.L().Error("ES init fiald", zap.Error(err))
		}
		zap.L().Debug(esversion)
		esclientV6[k] = client
	}
	return
}
