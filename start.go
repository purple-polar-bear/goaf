package main

import (
	"fmt"
	"log"
	"net/http"
	"oaf-server/codegen"
	"oaf-server/core"
	"oaf-server/geopackage"
	"github.com/purple-polar-bear/go-ogc-api/core"
	"github.com/purple-polar-bear/go-ogc-api/core/models"
	"github.com/purple-polar-bear/go-ogc-api/core/services"
	"github.com/purple-polar-bear/go-ogc-api/features"

	// ogcapi "oaf-server/package"
	// ogcapifeatures "oaf-server/package/features"

	"oaf-server/postgis"
	"oaf-server/server"
	"os"
	"regexp"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/urfave/cli/v2"

	"github.com/rs/cors"
)

func main() {

	app := cli.NewApp()
	app.Name = "GOAF"
	app.Usage = "A Golang OGC API Features implementation"
	app.HideVersion = true
	app.HideHelp = true

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "host",
			Aliases:     []string{"h"},
			Usage:       "server internal bind host address",
			DefaultText: "0.0.0.0",
			Required:    false,
			EnvVars:     []string{"HOST"},
		},
		&cli.StringFlag{
			Name:        "port",
			Aliases:     []string{"p"},
			Usage:       "server internal bind port",
			DefaultText: "8080",
			Required:    false,
			EnvVars:     []string{"PORT"},
		},
		&cli.StringFlag{
			Name:     "config",
			Aliases:  []string{"c"},
			Usage:    "Path to the configuration",
			Required: true,
			EnvVars:  []string{"CONFIG"},
		},
	}

	app.Action = func(c *cli.Context) error {

		configfilepath := c.String("config")
		config := &core.Config{}
		config.ReadConfig(configfilepath)

		// stage 1: create server with spec path and limits
		apiServer, err := server.NewServer(config.Service.Url, config.Openapi, uint64(config.DefaultFeatureLimit), uint64(config.MaxFeatureLimit))
		if err != nil {
			log.Fatal("Server initialisation error:", err)
		}

		// stage 2: Create providers based upon provider name
		// commonProvider := core.NewCommonProvider(config.Openapi, uint64(config.DefaultFeatureLimit), uint64(config.MaxFeatureLimit))
		providers := getProvider(apiServer.Openapi, *config)

		if providers == nil {
			log.Fatal("Incorrect provider provided valid names are: gpkg, postgis")
		}

		// stage 3: Add providers, also initialises them
		apiServer, err = apiServer.SetProviders(providers)
		if err != nil {
			log.Fatal("Server initialisation error:", err)
		}

		// stage 4: Prepare routing
		router := apiServer.Router()

		// extra routing for healthcheck
		addHealthHandler(router)

		// extra routing for package calls
		addPackageHandler(router, config)

		fs := http.FileServer(http.Dir("swagger-ui"))
		router.Handler(regexp.MustCompile("/swagger-ui"), http.StripPrefix("/swagger-ui/", fs))

		// cors handler
		handler := cors.Default().Handler(router)

		host := c.String("host")
		port := c.String("port")

		bindAddress := "0.0.0.0:8080"
		if host != "" || port != "" {
			if host == "" {
				host = "0.0.0.0"
			}

			if port == "" {
				port = "8080"
			}

			bindAddress = fmt.Sprintf("%v:%v", host, port)
		}

		// ServerEndpoint can be different from bind address due to routing externally
		log.Print("|\n")
		log.Printf("| SERVING ON: %s \n", apiServer.ServiceEndpoint)

		// stage 5: Start server
		return http.ListenAndServe(bindAddress, handler)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getProvider(api *openapi3.T, config core.Config) codegen.Providers {
	if config.Datasource.Geopackage != nil {
		return geopackage.NewGeopackageWithCommonProvider(api, config)
	} else if config.Datasource.PostGIS != nil {
		return postgis.NewPostgisWithCommonProvider(api, config)
	}

	return nil
}

func addHealthHandler(router *server.RegexpHandler) {
	router.HandleFunc(regexp.MustCompile("/health"), func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err := w.Write([]byte("ok"))
		if err != nil {
			log.Printf("Could not write ok")
		}
	})
}

//
// End of featureservice
//

func addPackageHandler(router *server.RegexpHandler, dsrc *core.Config) {
	mountingPath := "/package"

	// engine := ogcapi.NewSimpleEngine(mountingPath)
	// ogcapifeatures.EnableFeatures(engine, featuredatasource)

	engine := apicore.NewSimpleEngine(mountingPath)
	apicore.AddBaseJSONTemplates(engine)
	apicore.AddBaseHTMLTemplates(engine)

	config := engine.Config()
	config.SetTitle("goaf Demo instance - running latest GitHub version")
	config.SetDescription("goaf provides an API to geospatial data")
	config.SetProtocol("http")
	config.SetHost("172.17.0.1")
	config.SetPort(8080)

	apiService := engine.GetService("core").(coreservices.CoreService)
	apiService.SetContact(&coreservices.ContactInfo{Name: "PDOK", Url: "https://pdok.nl/contact"})
	apiService.SetLicense(&coreservices.LicenseInfo{Name: "CC-BY 4.0 license", Url: "https://creativecommons.org/licenses/by/4.0/"})
	ctUrlEncoder := coremodels.NewContentTypeUrlEncoding("f")
	ctUrlEncoder.AddContentType("json", "application/vnd.oai.openapi+json;version=3.0")
	ctUrlEncoder.AddContentType("json", "application/json")
	ctUrlEncoder.AddContentType("html", "text/html")
	apiService.SetContentTypeUrlEncoder(ctUrlEncoder)

	featuredatasource := geopackage.Init(*dsrc)
	apifeatures.EnableFeatures(engine, featuredatasource)
	apifeatures.AddFeaturesJSONTemplates(engine)
	apifeatures.AddFeaturesHTMLTemplates(engine)

	engine.RebuildOpenAPI()
	router.HandleFunc(regexp.MustCompile("^"+mountingPath), engine.HTTPHandler)
}
