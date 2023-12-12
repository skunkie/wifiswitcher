package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

// App struct
type App struct {
	ctx          context.Context
	cfg          config
	clientConfig *ssh.ClientConfig
	logger       *log.Logger
}

type config struct {
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// NewApp creates a new App application struct
func NewApp(cfg config, logger *log.Logger) *App {
	return &App{
		cfg: cfg,
		clientConfig: &ssh.ClientConfig{
			User: cfg.Username,
			Auth: []ssh.AuthMethod{
				ssh.Password(cfg.Password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		},
		logger: logger,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Toggle(upDown string) {
	var cmd string

	switch upDown {
	case "up":
		cmd = "/sbin/wifi up"
	case "down":
		cmd = "/sbin/wifi down"
	default:
		a.logger.Printf("unknown action %s\n", upDown)
		return
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port), a.clientConfig)
	if err != nil {
		a.logger.Printf("failed to dial, %s\n", err)
		return
	}
	session, err := client.NewSession()
	if err != nil {
		a.logger.Printf("failed to create session, %s\n", err)
		return
	}
	defer session.Close()

	if err := session.Run(cmd); err != nil {
		a.logger.Printf("failed to run command %q, %s\n", cmd, err.Error())
	}
}
