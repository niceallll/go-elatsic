//package cos
//
//import (
//	"context"
//	"fmt"
//	"go.uber.org/zap"
//	"time"
//	setting "webesapp/settings"
//
//	//"fmt"
//	"github.com/tencentyun/cos-go-sdk-v5"
//	"net/http"
//	"net/url"
//	//"os"
//	//"strings"
//)
//
//var c *cos.Client
//
//func Init(cfg *setting.Cos) (err error) {
//	u, _ := url.Parse(cfg.CosUrl)
//	b := &cos.BaseURL{BucketURL: u}
//	c := cos.NewClient(b, &http.Client{
//		Transport: &cos.AuthorizationTransport{
//			SecretID:     cfg.SecretID,  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
//			SecretKey:    cfg.SecretKey, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
//			SessionToken: "SECRETTOKEN",
//		},
//	})
//	fmt.Println(c.Host)
//	return
//}
//
//func Upload(filename string) (urls *url.URL) {
//	ak := "xx"
//	sk := "xx"
//	//key := "\\webapps\\" + filename
//	key := "/webapps/" + filename
//	fmt.Println(filename)
//	_, _, err := c.Object.Upload(
//		context.Background(), key, filename, nil,
//	)
//
//	if err != nil {
//		zap.L().Error("上传失败", zap.Error(err))
//	}
//	ctx := context.Background()
//	ourl, err := c.Object.GetPresignedURL(ctx, http.MethodGet, key, ak, sk, time.Hour, nil)
//
//	//ourl := c.Object.GetObjectURL(key)
//	fmt.Println(ourl)
//	if err != nil {
//		panic(err)
//	}
//
//	return ourl
//}

package cos

import (
	"context"
	"go.uber.org/zap"
	"time"

	//"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	//"os"
	//"strings"
)

func Upload(filename string) (urls *url.URL) {
	u, _ := url.Parse("https://xxxxxx.cos.ap-shanghai.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:     "xx", // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey:    "xx", // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SessionToken: "SECRETTOKEN",
		},
	})
	ak := "xxx"
	sk := "xxx"
	//key := "\\webapps\\" + filename
	key := "/webapps/" + filename
	_, _, err := c.Object.Upload(
		context.Background(), key, filename, nil,
	)

	if err != nil {
		zap.L().Error("上传失败", zap.Error(err))
	}
	ctx := context.Background()
	ourl, err := c.Object.GetPresignedURL(ctx, http.MethodGet, key, ak, sk, time.Hour, nil)

	//ourl := c.Object.GetObjectURL(key)
	if err != nil {
		panic(err)
	}

	return ourl
}
