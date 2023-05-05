/* !!
 * File: utils.go
 * File Created: Monday, 21st June 2021 3:26:45 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Monday, 21st June 2021 3:26:45 pm
 
 */

package g_api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/TrHung-297/fountain/baselib/base"
	"github.com/TrHung-297/fountain/baselib/crypto"
	"github.com/TrHung-297/fountain/proto/g_proto"
)

func ping(c *fiber.Ctx) error {
	return c.JSON(map[string]interface{}{
		"Message": "Success",
		"Data": map[string]interface{}{
			"message": "pong",
			"time":    time.Now().UTC().Format("2006-01-02 15:04:05"),
			"status":  "oke",
			"from":    c.Context().RemoteAddr().String(),
			"agent":   string(c.Context().UserAgent()),
		},
		"Status": 200,
	})
}

func config(c *fiber.Ctx) error {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		return WriteError(c, g_proto.NewRpcError(int32(g_proto.GTVRpcErrorCodes_ERROR_INTERNAL), err.Error()))
	}

	config := make(map[string]interface{})
	err = viper.Unmarshal(&config)
	if err != nil {
		return WriteError(c, g_proto.NewRpcError(int32(g_proto.GTVRpcErrorCodes_ERROR_INTERNAL), err.Error()))
	}

	cryptor := crypto.NewAES256IGECryptor([]byte("GAPIConfigTester"), []byte("GAPIConfigTester"))
	data, err := cryptor.EncryptSimpleWithBase64(base.JSONDebugData(config))
	if err != nil {
		return WriteError(c, g_proto.NewRpcError(int32(g_proto.GTVRpcErrorCodes_ERROR_INTERNAL), err.Error()))
	}

	cf, err := cryptor.DecryptSimpleWithBase64(string(data))
	if err != nil {
		return WriteError(c, g_proto.NewRpcError(int32(g_proto.GTVRpcErrorCodes_ERROR_INTERNAL), err.Error()))
	}
	return WriteSuccess(c, cf)
}

func GetContextDataString(ctx *fiber.Ctx, key string, defaultValues ...string) string {
	defaultValue := ""
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}

	userUUIDRaw := ctx.Locals(key)
	if userUUIDRaw != nil {
		if res, ok := userUUIDRaw.(string); ok {
			return res
		}
	}

	return defaultValue
}

// Parse data from param and set all into body
func ParamToBody(reqParams ...ReqParam) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if len(reqParams) == 0 {
			return c.Next()
		}

		data := make(map[string]interface{})
		if err := c.BodyParser(&data); err != nil {
			err := g_proto.NewRpcError(int32(g_proto.GTVRpcErrorCodes_ERROR_INTERNAL), err.Error())
			return WriteError(c, err)
		}

		for _, reqParam := range reqParams {
			data[reqParam.Name] = c.Params(reqParam.Param, reqParam.DefaultValue)
		}

		c.Request().SetBody(base.JSONDebugData(data))

		return c.Next()
	}
}

// Parse data from param and set all into body
func QueryToBody(reqParams ...ReqParam) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if len(reqParams) == 0 {
			return c.Next()
		}

		data := make(map[string]interface{})
		if err := c.BodyParser(&data); err != nil {
			err := g_proto.NewRpcError(int32(g_proto.GTVRpcErrorCodes_ERROR_INTERNAL), err.Error())
			return WriteError(c, err)
		}

		for _, reqParam := range reqParams {
			data[reqParam.Name] = c.Query(reqParam.Param, reqParam.DefaultValue)
		}

		c.Request().SetBody(base.JSONDebugData(data))

		return c.Next()
	}
}

// Parse data from param and set all into context
func ParamToContex(reqParams ...ReqParam) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if len(reqParams) == 0 {
			return c.Next()
		}

		for _, reqParam := range reqParams {
			c.Locals(reqParam.Name, c.Params(reqParam.Param, reqParam.DefaultValue))
		}

		return c.Next()
	}
}

// Parse data from query and set all into context
func QueryToContext(reqParams ...ReqParam) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if len(reqParams) == 0 {
			return c.Next()
		}

		for _, reqParam := range reqParams {
			c.Locals(reqParam.Name, c.Query(reqParam.Param, reqParam.DefaultValue))
		}

		return c.Next()
	}
}

// Parse all data from body and set all into context
func BodyToContext(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	if err := c.BodyParser(&data); err != nil {
		err := g_proto.NewRpcError(int32(g_proto.GTVRpcErrorCodes_ERROR_INTERNAL), err.Error())
		return WriteError(c, err)
	}

	for k, v := range data {
		c.Locals(k, v)
	}

	return c.Next()
}

// Set data{Name, DefaultValue} to context
func DataToContext(reqParams ...ReqParam) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if len(reqParams) == 0 {
			return c.Next()
		}

		for _, reqParam := range reqParams {
			c.Locals(reqParam.Name, reqParam.DefaultValue)
		}

		return c.Next()
	}
}
