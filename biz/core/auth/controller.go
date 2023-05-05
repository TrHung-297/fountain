

package auth

import (
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/firebase_auth"
	"gitlab.gplay.vn/gtv-backend/fountain/biz/core"
	"gitlab.gplay.vn/gtv-backend/fountain/biz/dal/dao"
	"gitlab.gplay.vn/gtv-backend/fountain/biz/dal/dao/redis_dao"
)

type authDAO struct {
	OpenIDCacheDAO *redis_dao.OpenIDCacheDAO
}

var authControllerInstance *AuthController

type AuthController struct {
	dao          *authDAO
	FirebaseAuth *firebase_auth.FirebaseAuth
}

func (ctrl *AuthController) InstallController() {
	ctrl.dao.OpenIDCacheDAO = dao.GetRedisOpenIDCacheDAO(dao.OPEN_ID_CACHE)

	authControllerInstance = ctrl
}

func (ctrl *AuthController) RegisterCallback(cb interface{}) {

}

func init() {
	core.RegisterCoreController(&AuthController{dao: &authDAO{}})
}

func GetAuthControllerInstance() *AuthController {
	return authControllerInstance
}

func (ctrl *AuthController) SetDAO(redisName string) {
	ctrl.dao.OpenIDCacheDAO = dao.GetRedisOpenIDCacheDAO(redisName)
}

func (ctrl *AuthController) InstallFirebaseAuth(configKeys ...string) {
	ctrl.FirebaseAuth = firebase_auth.InstallFirebaseAuth(configKeys...)
}
