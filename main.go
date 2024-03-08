package main

import (
	"fmt"
	"net/http"

	"os"
	"os/signal"
	"runtime"
	"syscall"

	"sampleProvider/cmd"
	"sampleProvider/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/urfave/cli"
)

const defaultPort = "8088"

func main() {
	if err := setupAPP().Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func setupAPP() *cli.App {
	app := cli.NewApp()
	app.Usage = "orange provider wrapper service"
	app.Action = startAgent
	app.Flags = []cli.Flag{
		cmd.LogLevelFlag,
		cmd.LogDirFlag,
		cmd.RpcUrlFlag,
		cmd.PortFlag,
		cmd.ConfigFileFlag,
		cmd.OperationFlag,
	}
	app.Before = func(context *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}
	return app
}

func startAgent(ctx *cli.Context) {
	port := ctx.GlobalString(cmd.GetFlagName(cmd.PortFlag))
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowedHeaders:   []string{"Authorization", "Content-Length", "X-CSRF-Token", "Token", "session", "X_Requested_With", "Accept", "Origin", "Host", "Connection", "Accept-Encoding", "Accept-Language", "DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Pragma"},
		ExposedHeaders:   []string{"Content-Length", "token", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Cache-Control", "Content-Language", "Content-Type", "Expires", "Last-Modified", "Pragma", "FooBar"},
		AllowCredentials: false,
		MaxAge:           172800, // Maximum value not ignored by any of major browsers
		// Debug:true,
	}))

	router.Route("/", func(r chi.Router) {
		r.Use(cors.Handler(cors.Options{
			// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins: []string{"*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			// AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowedHeaders:   []string{"Authorization", "Content-Length", "X-CSRF-Token", "Token", "session", "X_Requested_With", "Accept", "Origin", "Host", "Connection", "Accept-Encoding", "Accept-Language", "DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Pragma"},
			ExposedHeaders:   []string{"Content-Length", "token", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Cache-Control", "Content-Language", "Content-Type", "Expires", "Last-Modified", "Pragma", "FooBar"},
			AllowCredentials: false,
			MaxAge:           172800, // Maximum value not ignored by any of major browsers
			// Debug:true,
		}))
		r.Post("/sampleGetBodyDP", service.SampleGetBodyDP)
		r.Get("/sampleGetUrlDP", service.SampleGetGetUrlDP)

		r.Post("/sampleAP", service.SampleAP)

	})

	go signalHandle()
	fmt.Printf("staring restful at port:%s\n", port)
	http.ListenAndServe(":"+port, router)
}

func signalHandle() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Println("get a signal: stop the rest gateway process", si.String())
			os.Exit(1)
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
