package pkg

import (
	"github.com/cfi2017/wings-api/pkg/model"
)

type Api interface {
	RegisterHandler(name string, handler func(api Api))
	RegisterBackupStrategy(name string, strategy model.Backup)
	CallHandler(name string)
}
