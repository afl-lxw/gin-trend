package core

import (
	"fmt"
	"github.com/afl-lxw/gin-trend/global"
	"github.com/afl-lxw/gin-trend/initialize"
	"go.uber.org/zap"
	"moul.io/banner"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	if global.TREND_CONFIG.System.UseMultipoint || global.TREND_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}
	if global.TREND_CONFIG.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
	// 从db加载jwt数据
	//if global.TREND_DB != nil {
	//	system.LoadAll()
	//}
	Router := initialize.Routers()

	Router.Static("/form-generator", "./resource/page")

	//address := fmt.Sprintf(":%d", global.TREND_CONFIG.System.Addr)
	address := fmt.Sprintf("0.0.0.0:%d", global.TREND_CONFIG.System.Addr)
	s := initServer(address, Router)

	time.Sleep(10 * time.Microsecond)

	fmt.Println(banner.Inline("welcome to GO Trend"))
	global.TREND_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf(`默认自动化文档地址:%s/swagger/index.html /n开始运行 运行地址为%s 
	`,
		address, address)
	global.TREND_LOG.Error(s.ListenAndServe().Error())
}
