package asmkit

import (
	"github.com/op/go-logging"
)

const (
	Version = "0.0.2"
	Author  = "Yibin Wang"
	License = "BSD 3-Clause"
)

var log = logging.MustGetLogger("asmkit")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05} %{shortfunc} | %{level:.6s} %{color:reset} %{message}`,
)
