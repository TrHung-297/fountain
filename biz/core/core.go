

package core

import (
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"
)

// CoreModel interface
type CoreController interface {
	InstallController()
	RegisterCallback(cb interface{})
}

var Controllers = []CoreController{}

// RegisterCoreController func;
func RegisterCoreController(ctr CoreController) {
	Controllers = append(Controllers, ctr)
	g_log.V(1).Infof("RegisterCoreModel - Controller: %T", ctr)
}

// InstallCoreControllers func
// must be executed after the mysql / redis and other dependencies are installed
func InstallCoreControllers() []CoreController {
	g_log.V(3).Infof("InstallCoreControllers Num: %d", len(Controllers))

	for _, c := range Controllers {
		c.InstallController()
		for _, c2 := range Controllers {
			if c != c2 {
				c.RegisterCallback(c2)
			}
		}
	}

	return Controllers
}
