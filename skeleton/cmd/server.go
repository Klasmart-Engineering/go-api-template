// Copyright 2022 james.wilson. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/james.wilson/goapi/core/controller"
	"github.com/james.wilson/goapi/core/service"

	"github.com/drone/envsubst"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start api server",
	Run: func(cmd *cobra.Command, args []string) {
		configUnparsed, err := ioutil.ReadFile(config)

		if err != nil {
			panic(fmt.Sprintf(
				"Error while reading config file [%s]: %s",
				config,
				err.Error(),
			))
		}

		configParsed, err := envsubst.EvalEnv(string(configUnparsed))

		if err != nil {
			panic(fmt.Sprintf(
				"Error while parsing config file [%s]: %s",
				config,
				err.Error(),
			))
		}

		viper.SetConfigType("yaml")
		err = viper.ReadConfig(bytes.NewBuffer([]byte(configParsed)))

		if err != nil {
			panic(fmt.Sprintf(
				"Error while loading configs [%s]: %s",
				config,
				err.Error(),
			))
		}

		fs := service.NewFileSystem()

		if viper.GetString("app.log.output") != "stdout" {
			dir, _ := filepath.Split(viper.GetString("app.log.output"))

			if !fs.DirExists(dir) {
				if err := fs.EnsureDir(dir, 775); err != nil {
					panic(fmt.Sprintf(
						"Directory [%s] creation failed with error: %s",
						dir,
						err.Error(),
					))
				}
			}

			if !fs.FileExists(viper.GetString("app.log.output")) {
				f, err := os.Create(viper.GetString("app.log.output"))
				if err != nil {
					panic(fmt.Sprintf(
						"Error while creating log file [%s]: %s",
						viper.GetString("app.log.output"),
						err.Error(),
					))
				}
				defer f.Close()
			}
		}

		defaultLogger := middleware.DefaultLoggerConfig

		if viper.GetString("app.log.output") == "stdout" {
			log.SetOutput(os.Stdout)
			defaultLogger.Output = os.Stdout
		} else {
			f, _ := os.Create(viper.GetString("app.log.output"))
			log.SetOutput(f)
			defaultLogger.Output = f
		}

		lvl := strings.ToLower(viper.GetString("app.log.level"))
		level, err := log.ParseLevel(lvl)

		if err != nil {
			level = log.InfoLevel
		}

		log.SetLevel(level)

		if viper.GetString("app.log.format") == "json" {
			log.SetFormatter(&log.JSONFormatter{})
		} else {
			log.SetFormatter(&log.TextFormatter{})
		}

		viper.SetDefault("config", config)

		e := echo.New()

		if viper.GetString("app.mode") == "dev" {
			e.Debug = true
		}

		e.Use(middleware.LoggerWithConfig(defaultLogger))
		e.Use(middleware.RequestID())
		e.Use(middleware.BodyLimit("2M"))

		e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Timeout: time.Duration(viper.GetInt("app.timeout")) * time.Second,
		}))

		p := prometheus.NewPrometheus(viper.GetString("app.name"), nil)
		p.Use(e)

		e.GET("/favicon.ico", func(c echo.Context) error {
			return c.String(http.StatusNoContent, "")
		})

		// API GW Health
		e.GET("/_health", controller.Health)

		var runerr error

		if viper.GetBool("app.tls.status") {
			runerr = e.StartTLS(
				fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("app.port"))),
				viper.GetString("app.tls.crt_path"),
				viper.GetString("app.tls.key_path"),
			)
		} else {
			runerr = e.Start(
				fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("app.port"))),
			)
		}

		if runerr != nil && runerr != http.ErrServerClosed {
			panic(runerr.Error())
		}
	},
}

func init() {
	serverCmd.Flags().StringVarP(
		&config,
		"config",
		"c",
		"config.prod.yml",
		"Absolute path to config file (required)",
	)
	serverCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(serverCmd)
}
