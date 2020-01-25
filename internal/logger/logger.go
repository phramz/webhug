package logger

import (
	"github.com/google/logger"
	"io/ioutil"
)

func Logger() *logger.Logger {
	return logger.Init("webhug", true, false, ioutil.Discard)
}
