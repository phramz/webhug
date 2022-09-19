package logger

import (
	"io"

	"github.com/google/logger"
)

func Logger() *logger.Logger {
	return logger.Init("webhug", true, false, io.Discard)
}
