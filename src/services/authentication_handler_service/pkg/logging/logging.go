package logging

import (
	"os"

	core_logging "github.com/BlackspaceInc/BlackspacePlatform/src/libraries/core/core-logging/json"
)

type ServiceCore struct {
	Logger core_logging.ILog
}

var SvcCore = &ServiceCore{Logger: core_logging.JSONLogger}

func (s *ServiceCore) Fatal(err error, msg string){
	s.Logger.ErrorM(err, msg)
	os.Exit(1)
}

