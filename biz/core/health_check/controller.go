

package health_check

import (
	"github.com/TrHung-297/fountain/baselib/env"
	"github.com/TrHung-297/fountain/baselib/g_log"
	"github.com/TrHung-297/fountain/baselib/g_net/g_api"
	"github.com/TrHung-297/fountain/biz/core"
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
