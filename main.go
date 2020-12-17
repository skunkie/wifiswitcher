package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/lxn/walk"
	d "github.com/lxn/walk/declarative"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
)

type config struct {
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var c = &config{}

var logger *log.Logger

func exec(cmd string) {
	clientConfig := &ssh.ClientConfig{
		User: c.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port), clientConfig)
	if err != nil {
		logger.Printf("failed to dial, %s\n", err)
	}
	session, err := client.NewSession()
	if err != nil {
		logger.Printf("failed to create session, %s\n", err)
	}
	defer session.Close()

	if err := session.Run(cmd); err != nil {
		logger.Printf("failed to run command %q, %s\n", cmd, err.Error())
	}
}

func main() {

	logFile, err := os.OpenFile("wifiswitcher.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer logFile.Close()
	logger = log.New(logFile, "", log.LstdFlags)

	configFilepath := flag.String("c", "config.yml", "Path to config file")
	logVersion := flag.Bool("v", false, "Write version and commit to log file and exit")
	flag.Parse()

	if *logVersion {
		logger.Printf("Version: %s, Commit: %s\n", version, commit)
		os.Exit(0)
	}

	configFile, err := os.Open(*configFilepath)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Fatalf("create a config file %q\n", *configFilepath)
		}
		logger.Fatalf("error while reading the config file %q, %s\n", *configFilepath, err)
	}
	defer configFile.Close()

	if err := yaml.NewDecoder(configFile).Decode(c); err != nil {
		logger.Fatal(err)
	}

	icon, err := walk.NewIconFromResource("#2")
	if err != nil {
		log.Fatal(err)
	}

	d.MainWindow{
		Icon:   icon,
		Title:  "OpenWrt WiFi Switcher",
		Size:   d.Size{Width: 300, Height: 180},
		Layout: d.VBox{},
		Children: []d.Widget{
			d.HSplitter{
				Children: []d.Widget{
					d.PushButton{
						Text: "UP",
						OnClicked: func() {
							exec("/sbin/wifi up")
						},
					},
					d.PushButton{
						Text: "DOWN",
						OnClicked: func() {
							exec("/sbin/wifi down")
						},
					},
				},
			},
		},
	}.Run()
}
