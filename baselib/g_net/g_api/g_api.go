/* !!
 * File: g_api.go
 * File Created: Thursday, 27th May 2021 10:19:32 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 4:18:05 pm
 
 */

package g_api

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/env"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_net/g_etcd"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// GAPI type;
type GAPI struct {
	*fiber.App
	conf     *Config
	registry *g_etcd.GRegistry
	Store    *session.Store
}

var (
	instanceGAPI *GAPI
)

// NewGAPI func;
func NewGAPI(configs ...fiber.Config) *GAPI {
	if instanceGAPI != nil {
		return instanceGAPI
	}

	getGAPIConfigFromEnv()

	if GapiConfig == nil {
		err := fmt.Errorf("not found config for gtv api instance")
		panic(err)
	}
	a := fiber.Config{
		BodyLimit: 10 * 1024 * 1024, // this is the default limit of 4MB
	}
	fiberApp := fiber.New(configs...)
	fiberApp = fiber.New(a)
	fiberApp.Use(cors.New())
	fiberApp.Use(recover.New(recover.Config{EnableStackTrace: true}))
	g_log.Infof(`env.LogPrintLevel: %d - debugMode: %t`, env.LogPrintLevel, env.LogPrintLevel == 5)
	if env.LogPrintLevel == 5 {
		fiberApp.Use(logger.New(logger.Config{Next: func(c *fiber.Ctx) bool {
			return strings.Contains(c.Request().URI().String(), "healthcheck")
		}}))
	}

	fiberApp.Static("*", "./public")



	prefix := env.EndpointPrefix
	if prefix == "" {
		prefix = strings.TrimPrefix(env.ServiceName, "backend-")
		prefix = strings.Replace(prefix, "-", "/", 1)
	}
	prefix = strings.TrimSpace(prefix)

	gr := fiberApp.Group(prefix)

	g_log.V(1).Infof("NewGAPI - Ping path prefix: %s", prefix)

	gr.Get("/ping", ping)
	gr.Get("/config", config)

	fiberApp.Get("/ping", ping)

	var reg *g_etcd.GRegistry
	if GapiConfig.Discovery != nil {
		etcd := g_etcd.NewEtcdDiscovery(g_etcd.WithEtcdAddrs(GapiConfig.Discovery.Addrs...), g_etcd.WithEtcdUsername(GapiConfig.Discovery.Username), g_etcd.WithEtcdPassword(GapiConfig.Discovery.Password))
		reg = etcd.NewRegister(GapiConfig.Discovery)
	}

	key := "cookie:fountain-session" // Replace with your SESSION_SECRET or similar
	store := session.New(session.Config{
		KeyLookup:      key,
		Expiration:     30 * 24 * time.Hour,
		CookieHTTPOnly: true,
		CookieSecure:   false,
		CookiePath:     "/",
	})

	instanceGAPI = &GAPI{
		App:      fiberApp,
		conf:     GapiConfig,
		registry: reg,
		Store:    store,
	}

	return instanceGAPI
}

// GetGAPIInstance func;
func GetGAPIInstance() *GAPI {
	if instanceGAPI == nil {
		err := fmt.Errorf("you need init GTV api first")
		panic(err)
	}

	return instanceGAPI
}

func (api *GAPI) GetConfig() *Config {
	return api.conf
}

// Serve func;
// Need run in a goroutine
func (api *GAPI) Serve() {
	if GapiConfig.PprofEnabled {
		g_log.Infof("GAPI::Serve - CreatePProfServer..!")
		api.CreatePProfServer()
	}

	if api.registry != nil {
		go api.registry.Register()
	}

	if err := api.App.Listen(api.conf.Addr); err != nil {
		g_log.WithError(err).Errorf("GAPI::Serve - Listen Error: %+v", err)
	}
}

// Stop func;
func (api *GAPI) Stop() {
	if api.registry != nil {
		api.registry.Deregister()
	}

	time.AfterFunc(5*time.Second, func() {
		g_log.V(1).Infof("GAPI::Stop - Force by os.Exit")
		os.Exit(0)
	})

	if err := api.App.Shutdown(); err != nil {
		g_log.V(1).WithError(err).Errorf("GAPI::Stop - Error: %+v", err)
	}
}
