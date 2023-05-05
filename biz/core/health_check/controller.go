

package health_check

import (
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/env"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_net/g_api"
	"gitlab.gplay.vn/gtv-backend/fountain/biz/core"
)

type healthCheckDAO struct{}

type HealthCheckController struct {
	dao *healthCheckDAO
	api *g_api.GAPI
}

var healthCheckControllerInstance *HealthCheckController

func (ctrl *HealthCheckController) InstallController() {
	gr := ctrl.api.App.Group(env.EndpointPrefix)
	g_log.V(3).Infof("HealthCheckController::InstallController - prefix: %s", env.EndpointPrefix)

	gr.Get("/api/v2.0/healthcheck/status", ctrl.GetStatus)

	healthCheckControllerInstance = ctrl
}

func (ctrl *HealthCheckController) RegisterCallback(cb interface{}) {

}

func init() {
	core.RegisterCoreController(&HealthCheckController{dao: &healthCheckDAO{}, api: g_api.NewGAPI()})
}

func GetHealthCheckControllerInstance() *HealthCheckController {
	return healthCheckControllerInstance
}
