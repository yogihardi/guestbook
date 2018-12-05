package run

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/signal"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/yogihardi/guestbook/dao/boltdb"

	"github.com/yogihardi/guestbook/service"

	"github.com/yogihardi/guestbook/rest"
	"golang.org/x/net/context"

	"github.com/inconshreveable/log15"
	"github.com/urfave/cli"
)

var Command = cli.Command{
	Name:  "run",
	Usage: "Run the service",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "socket",
			Usage:  "REST API `socket` either as '[tcp://]<address>:<port>' or 'unix://<path>' string",
			EnvVar: "GUESTBOOK_SOCKET",
			Value:  "tcp://127.0.0.1:8080",
		},
	},
	Action: func(c *cli.Context) error {
		log := log15.New("module", "guestbook-service")

		var err error
		var listener net.Listener

		// create socket for API server
		socket := c.String("socket")
		if strings.HasPrefix(socket, "unix://") {
			f := strings.TrimPrefix(socket, "unix://")
			if _, err := os.Stat(f); err == nil {
				err = os.Remove(f)
				if err != nil {
					return err
				}
			}
			if listener, err = net.Listen("unix", f); err == nil {
				err = os.Chmod(f, 0770)
			}

		} else {
			if strings.HasPrefix(socket, "tcp://") {
				socket = strings.TrimPrefix(socket, "tcp://")
			}
			listener, err = net.Listen("tcp", socket)
		}
		if err != nil {
			return err
		}

		// capture interrupt signals from OS
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			s := <-sig
			fmt.Println()
			log.Info(fmt.Sprintf("signal %s received", s.String()))
			cancel()
		}()

		dir, err := ioutil.TempDir("", "guestbook")
		if err != nil {
			log.Error("failed to get temp dir", "err", err)
			return err
		}
		dbPath := dir + "/test-guestbook.boltdb"

		db, err := bolt.Open(dbPath, 0600, nil)
		if err != nil {
			log.Error("failed to open dbfile", "err", err)
			return err
		}

		boltdbctx, err := boltdb.NewBoltDBCtx(db)
		if err != nil {
			log.Error("failed to create boltdbctx", "err", err)
			return err
		}

		dao := boltdb.New(boltdbctx)

		// init application service
		appService, err := service.NewService(ctx, dao)
		if err != nil {
			log15.Error(err.Error())
			return err
		}

		return rest.Run(ctx, listener, appService, log)
	},
}
