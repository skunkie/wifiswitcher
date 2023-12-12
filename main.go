package main

import (
	"embed"
	"flag"
	"log"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"gopkg.in/yaml.v3"
)

const appName = "wifiswitcher"

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	var (
		cfgFilePath = flag.String("c", "config.yml", "Path to config file")
		logFilePath = flag.String("l", "wifiswitcher.log", "Path to log file")
	)
	flag.Parse()

	cfgFile, err := os.Open(*cfgFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("create a config file %q\n", *cfgFilePath)
		}
		log.Fatalf("error while reading the config file %q, %s\n", *cfgFilePath, err)
	}
	defer cfgFile.Close()

	logFile, err := os.OpenFile(*logFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("error while opening the log file %q, %s\n", *logFilePath, err)
	}
	defer logFile.Close()

	logger := log.New(logFile, appName+": ", log.LstdFlags)

	var cfg config

	if err := yaml.NewDecoder(cfgFile).Decode(&cfg); err != nil {
		logger.Fatal(err)
	}

	// Create an instance of the app structure
	app := NewApp(cfg, logger)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "OpenWrt WiFi Switcher",
		Width:  300,
		Height: 120,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		logger.Fatal(err)
	}
}
