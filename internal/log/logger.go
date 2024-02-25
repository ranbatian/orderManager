package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var Log *zap.Logger

func InitLogger() {
	c := zap.NewDevelopmentConfig()
	c.OutputPaths = append(c.OutputPaths, "./data/log/info.log")
	c.ErrorOutputPaths = append(c.ErrorOutputPaths, "./data/log/err.log")
	var err error
	Log, err = c.Build()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}

	zap.ReplaceGlobals(Log)

	defer Log.Sync()

}
