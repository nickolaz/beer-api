package logs

import (
	"fmt"
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func InitLogger() error {
	l, err := zap.NewDevelopment()
	// l , error := zap.NewProduction()

	if err != nil {
		_ = fmt.Errorf("Cannot create zap logger: %v", err)
		return err
	}

	sugar = l.Sugar()

	return nil
}

func Log() *zap.SugaredLogger {
	return sugar
}
