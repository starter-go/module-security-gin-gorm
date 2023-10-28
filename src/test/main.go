package main

import (
	"os"

	"github.com/starter-go/module-security-gin-gorm/modules/securitygingorm"
	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(os.Args)
	i.MainModule(securitygingorm.ModuleForTest())
	i.WithPanic(true).Run()
}
