package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	commonmodels "github.com/dashitoishi23/splitwise-go/models"
	db "github.com/dashitoishi23/splitwise-go/pkg/splits/database"
	splitendpoint "github.com/dashitoishi23/splitwise-go/pkg/splits/endpoints"
	splitservice "github.com/dashitoishi23/splitwise-go/pkg/splits/service"
	splitservers "github.com/dashitoishi23/splitwise-go/pkg/splits/transports"
	"github.com/dashitoishi23/splitwise-go/util"
	"github.com/joho/godotenv"
	"github.com/oklog/oklog/pkg/group"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println(err.Error())
	}

	var (
		httpAddr = fmt.Sprintf("0.0.0.0:%v", 9000)
	)

	db, dbErr := db.OpenDBConnection()

	var servers []commonmodels.HttpServerConfig

	if dbErr == nil {
		var (
			splitService  = splitservice.NewSplitService(db)
			splitEndpoint = splitendpoint.New(splitService)
			splitServers  = splitservers.NewHttpHandler(splitEndpoint)
		)

		servers = append(servers, splitServers...)

		httpHandler := util.RootHttpHandler(servers)

		httpListener, err := net.Listen("tcp", httpAddr)
		fmt.Println(httpListener.Addr().String(), err)

		var g group.Group
		{
			g.Add(func() error {
				fmt.Println(httpAddr)
				return http.Serve(httpListener, httpHandler)
			}, func(error) {
				httpListener.Close()
			})
		}
		{
			cancelInterrupt := make(chan struct{})
			g.Add(func() error {
				c := make(chan os.Signal, 1)
				signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
				select {
				case sig := <-c:
					return fmt.Errorf("received signal %s", sig)
				case <-cancelInterrupt:
					return nil
				}
			}, func(error) {
				close(cancelInterrupt)
			})
		}

		g.Run()

	}

	fmt.Println(dbErr.Error())

}
