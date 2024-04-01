package es

import (
	"fmt"
	//"fmt"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	setting "webesapp/settings"
)

var esclientV7 map[string]*elastic.Client

func InitV7(cfg *setting.ESconfig) (err error) {
	esclientV7 = make(map[string]*elastic.Client)
	for k, v := range cfg.Es7list {
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
		esclientV7[k] = client
	}
	return
}
