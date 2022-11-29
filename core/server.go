package core

import (
	"fmt"
	"time"

	"github.com/ciscolive/gin-admin-common/global"
	"github.com/ciscolive/gin-admin/initialize"
	"github.com/ciscolive/gin-admin/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.Config.System.UseMultipoint || global.Config.System.UseRedis {
		initialize.Redis()
	}

	if global.DB != nil {
		system.SyncJwtBlacklist()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.Config.System.Addr)
	s := initServer(address, Router)

	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	global.Logger.Info("成功加载项目", zap.String("后端监听地址和端口", address))
	fmt.Printf(`接口文档路径:http://127.0.0.1%s/swagger/index.html
前端访问地址:http://127.0.0.1:8080`, address)
	global.Logger.Error(s.ListenAndServe().Error())
}
