package service

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	CLI_USAGE       = ""
	CLI_DESCRIPTION = ""
)

// populated by `make build`
var (
	SERVICE_NAME    = ""
	GIT_COMMIT_ID   = ""
	GIT_BRANCH      = ""
	BUILD_TIMESTAMP = ""
)

func StartService() {

	fmt.Printf("Variables: %v %v %v", GIT_COMMIT_ID, GIT_BRANCH, BUILD_TIMESTAMP)

	cliApp := cli.NewApp()
	cliApp.Name = SERVICE_NAME
	cliApp.Usage = CLI_USAGE
	cliApp.Description = CLI_DESCRIPTION
	cliApp.Version = fmt.Sprintf("%v-%v (%v)", GIT_COMMIT_ID, GIT_BRANCH, BUILD_TIMESTAMP)

	cliApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "<service name> start",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "example-flag",
					EnvVar: "EXAMPLE_FLAG_ENV_CONSTANT",
					Value:  "example-flag-default-value",
				},
				cli.StringFlag{
					Name:        "retry-duration",
					Usage:       "",
					EnvVar:      "RETRY_DURATION",
					Required:    false,
					Hidden:      false,
					TakesFile:   false,
					Value:       "3000ms",
					Destination: nil,
				},
			},
			Action: handleServiceStart,
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func handleServiceStart(cliCtx *cli.Context) error {

	retryDelay, err := time.ParseDuration(cliCtx.String("retry-duration"))
	if err != nil {
		return err
	}

	log.Info("Starting %v %v-%v", SERVICE_NAME, GIT_COMMIT_ID, GIT_BRANCH)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//TODO: goroutines of the service have to be started here and they should be able to gracefully stopped, cancel functions should be passed to handleShutdowGracefully

	go handleShutdownGracefully(cancel)

	for {

		//TODO: do business logic here

		if err == context.Canceled {
			return err
		}

		log.Errorf("retrying in %s. error: %s", cliCtx.String("retry-duration"), err)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(retryDelay):
		}
	}
}

func handleShutdownGracefully (cancel context.CancelFunc) {
	sCh := make(chan os.Signal)
	signal.Notify(sCh, os.Interrupt, syscall.SIGTERM)
	<-sCh

	log.Info("Shutdown requested, cleaning up first")

	//TODO: goroutines of the service should be gracefully stopped here

	cancel()
	log.Info("Shutdown complete")
}

