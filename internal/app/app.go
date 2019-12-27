package app

import (
	"fmt"
	"github.com/urfave/cli"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	gitSummary = ""
	gitBranch = ""
	buildStamp = ""
)

func InitCLIApp() {
	app := cli.NewApp()
	app.Name = APP_NAME
	app.Usage = APP_USAGE
	app.Description =  APP_DESCRIPTION
	app.Version = fmt.Sprintf("%v-%v (%v)", gitSummary, gitBranch, buildStamp)

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "<service name> start",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "example flag",
					EnvVar: "EXAMPLE_FLAG_ENV_CONSTANT",
					Value:  "example-flag-default-value",
				},
			},
			Action: serviceActionHandler,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func serviceActionHandler(c *cli.Context) error {

	log.Info("service running")

	//retryDelay, err := time.ParseDuration(c.String("retry-duration"))
	//if err != nil {
	//	return err
	//}
	//downlineInterval, err := time.ParseDuration(c.String("downline-interval"))
	//if err != nil {
	//	return err
	//}
	//
	//log.Printf("Starting %v %v-%v", appName, gitSummary, gitBranch)
	//startOpListener(c.Int("operational-port"))
	//
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//
	//downline := NewDownlineExecutor()
	//downline.Interval = downlineInterval
	//go func() {
	//	sCh := make(chan os.Signal)
	//	signal.Notify(sCh, os.Interrupt, syscall.SIGTERM)
	//	<-sCh
	//
	//	log.Info("Shutdown requested, cleaning up first")
	//	downline.Close()
	//	cancel()
	//	log.Info("Shutdown complete")
	//}()
	//
	//
	//for {
	//	err := Index(ctx, c, chainedHandler)
	//	if err == nil {
	//		return nil
	//	}
	//	if err == context.Canceled {
	//		return err
	//	}
	//
	//	log.Errorf("retrying in %s. error: %s", c.String("retry-duration"), err)
	//	select {
	//	case <-ctx.Done():
	//		return ctx.Err()
	//	case <-time.After(retryDelay):
	//	}
	//}

	return nil
}
