/* !!
 * File: pprof.go
 * File Created: Monday, 8th November 2021 9:20:27 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Monday, 8th November 2021 9:20:34 pm
 
 */

package g_api

import (
	"fmt"
	"net/http/pprof"
	runtime_pprof "runtime/pprof"
	"time"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/TrHung-297/fountain/baselib/env"
	"github.com/TrHung-297/fountain/baselib/g_log"
)

// CreatePProfServer func;
func (api *GAPI) CreatePProfServer() {
	name := fmt.Sprintf("This is prof for %s:%s:%s at %s", env.DCName, env.ServiceName, env.PodID, env.Environment)
	pprofRoute := api.Group(env.EndpointPrefix)

	pprofRoute.Get("/debug/pprof/cpu", func(ctx *fiber.Ctx) error {
		ctx.WriteString(name)
		return runtime_pprof.StartCPUProfile(ctx)
	})

	// Register pprof handlers
	pprofRoute.Get("/debug/pprof/", adaptor.HTTPHandlerFunc(pprof.Index))
	pprofRoute.Get("/debug/pprof/cmdline", adaptor.HTTPHandlerFunc(pprof.Cmdline))
	pprofRoute.Get("/debug/pprof/profile", adaptor.HTTPHandlerFunc(pprof.Profile))
	pprofRoute.Get("/debug/pprof/symbol", adaptor.HTTPHandlerFunc(pprof.Symbol))
	pprofRoute.Get("/debug/pprof/trace", adaptor.HTTPHandlerFunc(pprof.Trace))
	pprofRoute.Get("/debug/runtime/pprof/trace", func(ctx *fiber.Ctx) error {
		err := runtime_pprof.StartCPUProfile(ctx)
		g_log.V(1).WithError(err).Errorf("GAPI::CreatePProfServer - Error: %+v", err)
		time.Sleep(30 * time.Second)
		runtime_pprof.StopCPUProfile()
		return nil
	})
	pprofRoute.Get("/debug/runtime/pprof/trace/heap", func(ctx *fiber.Ctx) error {
		return runtime_pprof.WriteHeapProfile(ctx)
	})
}
