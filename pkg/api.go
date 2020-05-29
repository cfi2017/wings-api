package pkg

import (
	"github.com/cfi2017/wings-api/pkg/model"
)

type Api interface {
	RegisterHandler(name string, handler func())
	RegisterBackupStrategy(name string, strategy model.Backup)
	CallHandler(name string)
}
