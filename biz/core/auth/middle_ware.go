
package auth

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/TrHung-297/fountain/baselib/g_log"
	"github.com/TrHung-297/fountain/baselib/g_net/g_api"
	"github.com/TrHung-297/fountain/baselib/g_token"
	"github.com/TrHung-297/fountain/proto/g_proto"
)

func (ctrl *AuthController) BaseGateway(c *fiber.Ctx, shortLive bool) (err error) {
	accessUUID, accountUUID, username, deviceKind, deviceIP, bearToken, err := g_token.ExtractToken(c, shortLive)
	if err != nil {
		if ctrl.FirebaseAuth != nil && strings.Contains(err.Error(), "can not parse token") {
			user := ctrl.FirebaseAuth.GetUser(bearToken)
			if user != nil {
				tokenErr := g_proto.NewRpcError(int32(g_proto.GTVRpcErrorCodes_ERROR_NEED_NEW_TOKEN), err.Error())
				g_log.V(3).Errorf("AuthController::JwtGateway - ExtractToken error: %+v", tokenErr)

				return g_api.WriteError(c, tokenErr)
			}
		}

		tokenErr := g_proto.NewRpcError(int32(g_proto.GTVRpcErrorCodes_ERROR_TOKEN_INVALID), err.Error())
		g_log.V(3).Errorf("AuthController::JwtGateway - ExtractToken error: %+v", tokenErr)

		return g_api.WriteError(c, tokenErr)
	}

	_ = deviceIP
	// REVIEW:
	// realIP := c.IP()
	// if deviceIP == "" && deviceIP != realIP {
	// 	err = fmt.Errorf("invalid Bearer Token - Localtion was changed from %s to %s", deviceIP, realIP)

	// 	tokenErr := g_proto.NewRpcError(int32(g_proto.GTVRpcErrorCodes_ERROR_TOKEN_INVALID), err.Error())
	// 	g_log.V(3).Errorf("AuthController::JwtGateway - ExtractToken error: %+v", tokenErr)

	// 	return g_api.WriteError(c, tokenErr)
	// }

	// Was user banend?

	// Compare REDIS
	var accessIDRedis string
	sessionKey := fmt.Sprintf(g_proto.KCacheUserAuthorizationFormat, accountUUID)

	if accessIDRedis = ctrl.dao.OpenIDCacheDAO.GetAccessIDByDeviceKind(sessionKey, deviceKind); accessIDRedis == "" || accessIDRedis != accessUUID {
		tokenErr := g_proto.NewRpcError(int32(g_proto.GTVRpcErrorCodes_ERROR_TOKEN_INVALID), fmt.Sprintf("invalid Bearer Token - can not compare token: %s vs %s in the store for user: %s with deviceKind: %d", accessUUID, accessIDRedis, accountUUID, deviceKind))
		g_log.V(3).Errorf("AuthController::JwtGateway - ExtractToken error: %+v", tokenErr)

		return g_api.WriteError(c, tokenErr)
	}

	c.Locals("AccessUUID", accessUUID)
	c.Locals("DeviceKind", deviceKind)
	c.Locals("DeviceIP", deviceIP)
	c.Locals("UserID", accountUUID)
	c.Locals("UserName", username)

	return c.Next()
}

func (ctrl *AuthController) FirebaseAuthGateway(c *fiber.Ctx) (err error) {
	_, bearToken, _ := g_token.GetToken(c)
	if bearToken == "" {
		tokenErr := g_proto.NewRpcError(int32(g_proto.GTVRpcErrorCodes_ERROR_TOKEN_INVALID), err.Error())
		g_log.V(3).Errorf("AuthController::JwtGateway - ExtractToken error: %+v", tokenErr)

		return g_api.WriteError(c, tokenErr)
	}

	c.Locals("BearToken", bearToken)

	return c.Next()
}

func (ctrl *AuthController) JwtGateway(c *fiber.Ctx) (err error) {
	return ctrl.BaseGateway(c, false)
}

func (ctrl *AuthController) SecureGateway(c *fiber.Ctx) (err error) {
	return ctrl.BaseGateway(c, true)
}
