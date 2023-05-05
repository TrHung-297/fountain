/* !!
 * File: resp_write.go
 * File Created: Wednesday, 26th May 2021 11:41:46 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Wednesday, 26th May 2021 11:41:46 am
 
 */

package g_api

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/base"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"
	"gitlab.gplay.vn/gtv-backend/fountain/proto/g_proto"
	"gopkg.in/go-playground/validator.v9"
)

/**
 * Defines a response object
 */
type Response struct {
	Message    string                                       `json:"Message,omitempty"`
	Data       interface{}                                  `json:"Data,omitempty"`
	DataExtend map[g_proto.GTVConstructor]g_proto.GTVObject `json:"DataExtend,omitempty"`
	Status     int                                          `json:"Status,omitempty"`
}

func NewGErrorResponse(err *g_proto.GTVRpcError) (string, *g_proto.GTVRpcError) {
	return err.GetMessage(), err
}

// func (Game) TableName() string { return "game" }

/**
 * Returns a success response
 */
func WriteSuccess(c *fiber.Ctx, v interface{}) error {
	res := Response{
		Message: "Success",
		Data:    v,
		Status:  200,
	}

	// Log response
	g_log.V(5).WithField("Route", c.Request().URI().String()).WithField("Response", base.JSONDebugDataString(res)).Info()

	// Return
	return c.JSON(res)
}

/**
 * Returns a success response without content
 */
func WriteSuccessEmptyContent(c *fiber.Ctx) error {
	res := Response{
		Message: "Success",
		Status:  200,
	}

	// Return
	return c.JSON(res)
}

/**
 * Returns a success response
 */
func WriteSuccessWithExtend(c *fiber.Ctx, data interface{}, extend map[g_proto.GTVConstructor]g_proto.GTVObject) error {
	res := Response{
		Message:    "Success",
		Data:       data,
		DataExtend: extend,
		Status:     200,
	}

	// Log response
	g_log.V(5).WithField("Route", c.Request().URI().String()).WithField("Response", base.JSONDebugDataString(res)).Info()

	// Return
	return c.JSON(res)
}

/**
 * Returns an error response
 */
func WriteError(c *fiber.Ctx, err *g_proto.GTVRpcError) error {
	res := Response{
		Message: "Error",
		Data:    err,
		Status:  int(err.Code),
	}

	// Log response
	g_log.V(5).WithField("Route", c.Request().URI().String()).WithError(err).WithField("Response", base.JSONDebugDataString(res)).Info()

	// Return
	return c.Status(int(err.Code)).JSON(res)
}

/**
 * Returns a success response
 */
func WriteErrorWithExtend(c *fiber.Ctx, err *g_proto.GTVRpcError, extend map[g_proto.GTVConstructor]g_proto.GTVObject) error {
	res := Response{
		Message:    "Error",
		Data:       err,
		DataExtend: extend,
		Status:     int(err.Code),
	}

	// Log response
	g_log.V(5).WithField("Route", c.Request().URI().String()).WithError(err).WithField("Response", base.JSONDebugDataString(res)).Info()

	// Return
	return c.JSON(res)
}

/**
 * Validates model before do something
 */
func IsValid(m interface{}) (bool, error) {
	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
