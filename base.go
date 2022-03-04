package asmkit

import (
	"github.com/op/go-logging"
)

const (
	Version = "0.0.1"
)

var log = logging.MustGetLogger("asmkit")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05} %{shortfunc} | %{level:.6s} %{color:reset} %{message}`,
)
