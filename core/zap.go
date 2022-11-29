package core

import (
	"fmt"
	"os"

	"github.com/ciscolive/gin-admin-common/core/internal"
	"github.com/ciscolive/gin-admin-common/global"
	"github.com/ciscolive/gin-admin/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.Config.Zap.Directory); !ok { // 判断是否有Directory文件夹
		// 如果不存在定义的文件夹则自动创建
		fmt.Printf("create %v directory\n", global.Config.Zap.Directory)
		_ = os.Mkdir(global.Config.Zap.Directory, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	// 是否加载调用者信息
	if global.Config.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
