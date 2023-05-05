

package health_check

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_net/g_api"
)

/**
 * Returns status
 */
func (ctrl *HealthCheckController) GetStatus(c *fiber.Ctx) error {
	return g_api.WriteSuccessEmptyContent(c)
}
