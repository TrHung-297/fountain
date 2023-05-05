package fountain_v2

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"github.com/TrHung-297/fountain/baselib/env"
	"github.com/TrHung-297/fountain/baselib/g_log"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
)

var GAppInstance AppInstance

// var enablePProf bool

func init() {
	flag.Parse()

	viper.SetConfigFile(`config.json`)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	env.SetupConfigEnv()

	logPath := env.LogPath
	logPrefix := env.LogPrefix

	if viper.GetBool(`Debug`) || env.LogPrintLevel == 5 {
		log.Println("Service RUN on DEBUG mode")
	} else {
		log.Println("Service RUN on PRODUCTION mode")
	}

	if logPath == "" {
		logPath = "/var/log/GtvPlus"
	}

	if logPrefix == "" {
		logPrefix = "backend-game"
	}

	g_log.LogDir(logPath)
	g_log.EnableTracing(false)
	g_log.LogFileLevel(uint32(env.LogFileLevel))
	g_log.LogPrintLevel(uint32(env.LogPrintLevel))

	g_log.Infof("logPath: %s", logPath)
	g_log.Infof("logPrefix: %s", logPrefix)
}

// AppInstance type;
type AppInstance interface {
	Initialize() error
	RunLoop()
	Destroy()
	GetIdentification() (addr, dcName, serviceName, serverID string)
}

var ch = make(chan os.Signal, 1)

// DoMainAppInstance func;
func DoMainAppInstance(instance AppInstance) {
	// Init sentry
	if sentryDNS := viper.GetString("Sentry.Dsn"); sentryDNS != "" {
		sentrySyncTransport := sentry.NewHTTPSyncTransport()
		sentrySyncTransport.Timeout = time.Second * 5

		err1 := sentry.Init(sentry.ClientOptions{
			Dsn:       sentryDNS,
			Transport: sentrySyncTransport,
		})
		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetTags(map[string]string{"module": "Log-Events-API"})
		})

		if err1 != nil {
			g_log.V(1).Infof("Sentry initialization failed: %v", err1)
		}

		sentry.CaptureMessage(fmt.Sprintf("Serice %s is running", os.Getenv("PODNAME")))

		// Flush buffered events before the program terminates.
		defer sentry.Flush(2 * time.Second)
	}

	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	// Recovery error handler
	defer func() {
		if f := recover(); f != nil {
			if err, ok := f.(error); ok {
				sentry.CaptureException(err)
				g_log.V(1).Infof("Global App Recover Error: %+v", err)
				debug.PrintStack()
			}
		}
	}()

	// if runtime.NumCPU() > 1 {
	// 	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	// }

	// Seeding for random
	rand.Seed(time.Now().UnixNano())

	if instance == nil {
		err := fmt.Errorf("instance is nil")
		g_log.V(1).Errorf("Error: %+v, will exit!", err)

		panic(err)
	}

	// Set global instance
	GAppInstance = instance

	_, _, serviceName, serverID := instance.GetIdentification()
	g_log.ServiceName(serviceName)
	g_log.ServerID(serverID)

	g_log.V(1).Info("Instance initialize...")
	err := instance.Initialize()
	if err != nil {
		g_log.V(1).Infof("instance initialize error: {%v}", err)
		panic(err)
	}

	// Profile instance debug
	// addr, dcName, serverName, serverID := instance.GetIdentification()
	// if enablePProf {
	// remotelog client
	// remotelog.InstallRemoteLogClient()
	// go pprof_trace.CreatePProfServer(addr, dcName, serverName, serverID)
	// }

	// Remote log

	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)

	g_log.V(1).Info("Instance run_loop...")
	go instance.RunLoop()

	g_log.V(1).Info("Wait quit...")

	s2 := <-ch
	if i, ok := s2.(syscall.Signal); ok {
		g_log.V(1).Infof("instance recv os.Exit(%d) signal...", i)
	} else {
		g_log.V(1).Infof("instance exit... with code %d", i)
	}

	instance.Destroy()

	g_log.V(1).Infof("instance quited!")

	time.Sleep(1 * time.Second)
	os.Exit(0)
}

// QuitAppInstance func;
func QuitAppInstance() {
	ch <- syscall.SIGQUIT
}
