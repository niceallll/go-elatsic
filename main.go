package main

import (
	"fmt"
	"go.uber.org/zap"
	"webesapp/dao/es"
	"webesapp/dao/mysql"
	"webesapp/dao/sqlserver"
	"webesapp/logger"
	"webesapp/routes"
	"webesapp/settings"
)

func main() {
	// 1.加载配置文件
	if err := setting.Init(); err != nil {
		fmt.Printf("init settings failed,err  %v\n", err)
		zap.L().Debug("settings init susses")
		return
	}

	// 2.初始化日志
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed,err  %v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init susses")

	// 3.初始化数据库sqlserver
	if err := sqlserver.Init(setting.Conf.SQLserverConfig); err != nil {
		fmt.Printf("init sqlser failed ,err %v\n", err)
		zap.L().Debug("sqlserver init failed")
	}
	zap.L().Debug("sqlserver init susses")
	defer sqlserver.Close()

	// 3.初始化数据库mysql
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed,err  %v\n", err)
		zap.L().Debug("mysql init failed")
		return
	}
	zap.L().Debug("mysql init susses")
	defer mysql.Close()

	//// 初始化es连接
	if err := es.InitV6(setting.Eslistconfig); err != nil {
		fmt.Printf("init es6 failed,err %v\n", err)
		zap.L().Debug("es6 init failed")
		return
	}
	zap.L().Debug("es6 init susses")

	//初始化es7连接
	if err := es.InitV7(setting.Eslistconfig); err != nil {
		fmt.Printf("init es7 failed,err %v\n", err)
		zap.L().Debug("es7 init failed")
		return
	}
	zap.L().Debug("es7 init susses")

	//if err := cos.Init(setting.Conf.Cos); err != nil {
	//	fmt.Printf("init cos failed,err %v\n", err)
	//	zap.L().Debug("cos init failed")
	//	return
	//}
	//zap.L().Debug("cos init susses")
	//test
	//if err := es.InitV7(setting.Eslistconfig); err != nil {
	//	fmt.Printf("init es failed,err %v\n", err)
	//	zap.L().Debug("es init susses")
	//	return
	//}
	//zap.L().Debug("es init susses")
	//
	//if err := es.Init3(setting.Eslistconfig); err != nil {
	//	fmt.Printf("init es failed,err %v\n", err)
	//	zap.L().Debug("es init susses")
	//	return
	//}
	zap.L().Debug("es init susses")
	// 4.注册路由
	r := router.SetupRouter(setting.Conf.Mode)
	// 5.启动服务
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

}
