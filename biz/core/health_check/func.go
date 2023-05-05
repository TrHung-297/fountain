

package health_check

import (
	"github.com/gofiber/fiber/v2"
	"github.com/TrHung-297/fountain/baselib/g_net/g_api"
)

/**
 * Returns status
 */
func (ctrl *HealthCheckController) GetStatus(c *fiber.Ctx) error {
	return g_api.WriteSuccessEmptyContent(c)
}
