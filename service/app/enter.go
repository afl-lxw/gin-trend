package app

import (
	"github.com/afl-lxw/gin-trend/service/app/auth"
	"github.com/afl-lxw/gin-trend/service/app/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	AuthServiceGroup   auth.ServiceGroup
}
