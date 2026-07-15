package main

import (
	"fmt"

	"github.com/acexy/golang-toolkit/logger"
	"github.com/acexy/golang-toolkit/sys"
	"github.com/golang-acexy/cloud-simple-demo/internal/handler/rest/adm"
	"github.com/golang-acexy/cloud-simple-demo/internal/handler/rest/usr"
	"github.com/golang-acexy/starter-gin/ginstarter"
	"github.com/golang-acexy/starter-gorm/gormstarter"
	"github.com/golang-acexy/starter-parent/parent"
)

var starterLoader *parent.StarterLoader

func init() {
	logger.EnableConsole(logger.TraceLevel, false)
	starterLoader = parent.InitStarterLoader([]parent.Starter{
		&gormstarter.GormStarter{
			Config: gormstarter.GormConfig{
				MySQL: &gormstarter.MySQLConfig{
					DatabaseConfig: gormstarter.DatabaseConfig{
						Username: "root",
						Password: "root",
						Database: "test",
						Host:     "127.0.0.1",
						Port:     13306,
					},
				},
			},
		},
		&ginstarter.GinStarter{
			Config: ginstarter.GinConfig{
				ListenAddress:     ":8080",
				UseReusePortModel: true,
				DebugModule:       true,
				Routers: []ginstarter.Router{
					usr.NewEmployeeRouter(),
					adm.NewDepartmentRouter(),
					adm.NewEmployeeRouter(),
				},
				TraceIDResponse: func() string {
					return logger.GetTraceId()
				},
			},
		},
	})

	err := starterLoader.Start()
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
}

func main() {
	sys.ShutdownHolding()
}
