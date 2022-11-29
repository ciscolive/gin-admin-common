package internal

import (
	"os"
	"path"
	"time"

	"github.com/ciscolive/gin-admin-common/global"
	logger "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

var RotateLog = new(rotateLog)

// 设定空接口
type rotateLog struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (r *rotateLog) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := logger.New(
		path.Join(global.Config.Zap.Directory, "%Y-%m-%d", level+".log"),        // 日志路径
		logger.WithClock(logger.Local),                                          // 日志文件锁
		logger.WithMaxAge(time.Duration(global.Config.Zap.MaxAge)*24*time.Hour), // 日志留存时间
		logger.WithRotationTime(time.Hour*24),                                   // 日志切割时间
	)
	// 是否同步打印日志到 console
	if global.Config.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
