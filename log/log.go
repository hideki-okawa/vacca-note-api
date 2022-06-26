package log

import "go.uber.org/zap"

var sugar *zap.SugaredLogger

func init() {
	logger, err := zap.NewDevelopment(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar = logger.Sugar()
}

// Debugf はdebugレベルのログを出力します。
func Debugf(template string, fields ...interface{}) {
	sugar.Debugf(template, fields...)
}

// Infof はinfoレベルのログを出力します。
func Infof(template string, fields ...interface{}) {
	sugar.Infof(template, fields...)
}

// Errorf はerrorレベルのログを出力します。
func Errorf(template string, fields ...interface{}) {
	sugar.Errorf(template, fields...)
}
